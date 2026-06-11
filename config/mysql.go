package config

import (
	"CWDLaunchPad/constants"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitliaseMySQLConnection() error {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv(constants.MySQLUserName), os.Getenv(constants.MySQLPassword), os.Getenv(constants.MySQLHost), os.Getenv(constants.MySQLDatabase)))
	if err != nil {
		return err
	}

	fmt.Println("DB connected successfully")

	return nil
}
