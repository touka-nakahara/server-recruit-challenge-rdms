package mysql

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mySqlDb struct {
	db *sql.DB
}

func NewMySQLDB() *mySqlDb {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	return &mySqlDb{db: db}
}

func connect() (*sql.DB, error) {
	//TODO いい感じにできるらしいが
	db, err := sql.Open("mysql", "root:root@tcp(:3307)/mysql")

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//TODO タイムアウトしたらコンテキストを切る
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil

}
