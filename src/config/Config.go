package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Environment() (string, error) {
	env := "development"

	return env, nil
}
func Hostname() (string, error) {
	host := "localhost"

	return host, nil
}

func DetermineListenAddresslinkaja() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		//return "", fmt.Errorf("$PORT not set")
		port = "2021"
	}
	return ":" + port, nil
}

const (
	DB_HOST = "tcp(localhost:3306)"
	DB_NAME = "linkajadb"
	DB_USER = "root"
	DB_PASS = ""
)

func Init() *sql.DB {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/linkajadb")
	checkErr(err)

	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)
	//fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
