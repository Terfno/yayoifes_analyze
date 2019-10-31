package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(127.0.0.1:63306)"
	DB_NAME := "yayoifes_analyze_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal("DB接続に失敗")
	}

	return db
}
