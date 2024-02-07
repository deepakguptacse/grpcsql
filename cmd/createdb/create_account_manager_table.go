package main

import (
	"database/sql"
	"log"

	"github.com/deepakguptacse/grpcsql/configs"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := configs.ReadConfig()
	if err != nil {
		log.Fatalf("Couldn't read config: %v", err)
	}
	db, err := sql.Open("mysql", cfg.SQLAddress)
	if err != nil {
		log.Fatalf("Couldn't connect to SQL server: %v", err)
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")
	database := dialect.DB(db)

	_, err = database.Exec("CREATE DATABASE IF NOT EXISTS grpcsql")
	if err != nil {
		log.Fatalf("Couldn't create database: %v", err)
	}

	_, err = database.Exec("USE grpcsql")
	if err != nil {
		log.Fatalf("Couldn't switch to database: %v", err)
	}

	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		id VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		active BOOLEAN NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		log.Fatalf("Couldn't create table: %v", err)
	}

	log.Println("Database grpcsql and table accounts created successfully")
}
