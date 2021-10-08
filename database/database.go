package database

import (
	"database/sql"
	"fmt"

	"github.com/b4cktr4ck5r3/nade404api/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config("user"),
		config.Config("password"),
		config.Config("host"),
		config.Config("port"),
		config.Config("dbname")))

	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}
