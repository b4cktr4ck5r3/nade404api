package database

import (
	"database/sql"
	"fmt"

	"github.com/b4cktr4ck5r3/nade404api/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectWithArgs(dbUser string, dbPwd string, dbHost string, dbPort string, dbName string) error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPwd,
		dbHost,
		dbPort,
		dbName))

	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}

func ConnectWithEnv() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config("DBUSER"),
		config.Config("DBPWD"),
		config.Config("DBHOST"),
		config.Config("DBPORT"),
		config.Config("DBNAME")))

	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}
