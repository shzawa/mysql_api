package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	// .env読み込み
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	// db接続
	db, err := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME"))
	if err != nil {
		panic(err.Error())
	}

	return db
}