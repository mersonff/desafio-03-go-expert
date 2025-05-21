package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	// Configuração do banco de dados
	dsn := "root:root@tcp(localhost:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
} 