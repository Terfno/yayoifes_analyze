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

func GetNumberOfVisitor() ([]*Log, error) {
	var l []*Log

	err := infra.Connect().Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}

func GetNumberOfVisitorByHour(min, max int) ([]*Log, error) {
	var l []*Log
	if min == 0 || max == 0 {
		err := infra.Connect().Find(&l, "Timing >= ? - INTERVAL 1 HOUR", time.Now()).Error
		if err != nil {
			return nil, err
		}
		return l, err
	}

	err := infra.Connect().Find(&l, "Timing BETWEEN (? - INTERVAL ? HOUR) AND (? - INTERVAL ? HOUR)", time.Now(), max, time.Now(), min).Error
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
	if err != nil {
		return nil, err
	}

	return l, err
}
