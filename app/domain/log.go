package domain

import (
	// "github.com/jinzhu/gorm"
	"time"

	"github.com/Terfno/yayoifes_analyze/app/infra"
)

type Log struct {
	Id     int       `json:id`
	Age    int       `json:age`
	Sex    string    `json:sex`
	Timing time.Time `json:datetime`
}

func GetTest() ([]*Log, error) {
	db := infra.Connect()
	defer db.Close()

	var l []*Log
	err := db.Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}
