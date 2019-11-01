package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Log struct {
	gorm.Model
	id     int       `json:id`
	age    int       `json:age`
	sex    string    `json:sex`
	timing time.Time `json:datetime`
}
