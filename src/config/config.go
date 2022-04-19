package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "vue_golang"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	fmt.Println("Successfully connected!", db)
	return

}
