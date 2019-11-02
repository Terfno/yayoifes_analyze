package domain

import (
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
	var l []*Log

	err := infra.Connect().Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}

func Read24() ([]*Log, error) {
	db := infra.Connect()
	defer db.Close()

	var l []*Log
	err := db.Find(&l, "Timing >= ? - INTERVAL 1 DAY", time.Now()).Error

	return l, err
}
