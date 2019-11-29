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

func NewReception(age int, sex string) error {
	db := infra.Connect()
	defer db.Close()
	newlog := Log{}
	newlog.Age = age
	newlog.Sex = sex
	newlog.Timing = time.Now()

	return db.Create(&newlog).Error
}

func GetNumberOfVisitor() ([]*Log, error) {
	var l []*Log

	db := infra.Connect()
	defer db.Close()

	err := db.Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}

func GetNumberOfVisitorByHour(min, max int) ([]*Log, error) {
	var l []*Log

	db := infra.Connect()
	defer db.Close()

	if min == 0 || max == 0 {
		err := db.Find(&l, "Timing >= ? - INTERVAL 1 HOUR", time.Now()).Error
		if err != nil {
			return nil, err
		}
		return l, err
	}

	err := db.Find(&l, "Timing BETWEEN (? - INTERVAL ? HOUR) AND (? - INTERVAL ? HOUR)", time.Now(), max, time.Now(), min).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func Read1103() ([]*Log, error) {
	var l []*Log

	db := infra.Connect()
	defer db.Close()

	err := db.Find(&l, "Timing between '2019-11-02 08:00:00' and '2019-11-03 08:00:00'").Error
	if err != nil {
		return nil, err
	}

	return l, err
}

func Read1104() ([]*Log, error) {
	var l []*Log

	db := infra.Connect()
	defer db.Close()

	err := db.Find(&l, "Timing >= '2019-11-03 08:00:00'").Error
	if err != nil {
		return nil, err
	}

	return l, err
}
