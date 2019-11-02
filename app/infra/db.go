package infra

import (
	// "time"
	"log"

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
		// for i := 0; i < 10; i++ {
		// 	db, err = gorm.Open(DBMS, CONNECT)
		// 	if err == nil {
		// 		break
		// 	}
		// 	time.Sleep(1 * time.Second)
		// }
		log.Fatal("die")
	}

	return db
}
