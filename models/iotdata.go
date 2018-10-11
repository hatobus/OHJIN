package models

import (
	"time"
)

type Iotdata struct {
	Co2         string    `xorm:"VARCHAR(10)"`
	Gettime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Humid       string    `xorm:"VARCHAR(10)"`
	Illuminance string    `xorm:"VARCHAR(10)"`
	Machineid   string    `xorm:"not null VARCHAR(3)"`
	No          int       `xorm:"not null pk autoincr INT(11)"`
	SoilHumid   string    `xorm:"VARCHAR(10)"`
	Temp        string    `xorm:"VARCHAR(10)"`
	Wavelength  string    `xorm:"VARCHAR(10)"`
}
