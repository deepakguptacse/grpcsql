package server

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"go.uber.org/zap"

	"github.com/deepakguptacse/grpcsql/configs"
	"github.com/deepakguptacse/grpcsql/proto"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type Server struct {
	proto.UnimplementedAccountManagerServer

	config *configs.Config
	db     *goqu.Database
}

func NewServer(cfg *configs.Config, db *goqu.Database) *Server {
	return &Server{config: cfg, db: db}
}

func validate(in *proto.CreateAccountRequest) error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, in.Email)
	if !match {
		return fmt.Errorf("invalid email format")
	}
	if in.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func (s *Server) CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	zap.S().Infof("Create account request received: %v", in)
	err := validate(in)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	query := `INSERT INTO accounts (id, name, email, active) VALUES (?, ?, ?, ?)`
	_, err = s.db.ExecContext(ctx, query, id, in.Name, in.Email, true)
	if err != nil {
		return nil, err
	}
	return &proto.CreateAccountResponse{Account: &proto.Account{Id: id}}, nil
}

func (s *Server) GetAccount(ctx context.Context, in *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	zap.S().Infof("Get account request received: %v", in)
	if in.Id == "" {
		return nil, fmt.Errorf("id is required")
	}
	query := `SELECT id, name, email, active FROM accounts WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, in.Id)
	account := &proto.Account{}
	err := row.Scan(&account.Id, &account.Name, &account.Email, &account.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}
	return &proto.GetAccountResponse{Account: account}, nil
}
