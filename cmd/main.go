package main

import (
	"database/sql"
	"log"
	"net"

	"go.uber.org/zap"

	"github.com/deepakguptacse/grpcsql/configs"
	"github.com/deepakguptacse/grpcsql/proto"
	"github.com/deepakguptacse/grpcsql/server"
	"github.com/doug-martin/goqu/v9"

	// "github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// TODO: use the production log in the prod env.
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)

	cfg, err := configs.ReadConfig()
	if err != nil {
		log.Fatalf("Couldn't read config: %v", err)
	}
	zap.S().Infof("SQL Address: %s", cfg.SQLAddress)

	// initialise the database connection
	db, err := sql.Open("mysql", cfg.SQLAddress)
	if err != nil {
		log.Fatalf("Couldn't connect to SQL server: %v", err)
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")
	database := dialect.DB(db)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAccountManagerServer(s, server.NewServer(cfg, database))
	reflection.Register(s)
	zap.S().Infof("Server started at port :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
