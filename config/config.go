package config

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var cfg = mysql.Config{
	User:   "root",
	Passwd: "admin",
	Net:    "tcp",
	Addr:   "localhost:3306",
	DBName: "abyshopdb",
}

// init database
func Init() {
	v := viper.New()

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	v.Unmarshal(db)
}
