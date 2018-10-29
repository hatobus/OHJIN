package models

import (
	"time"
)

type Iotdata struct {
	Co2         string    `xorm:"VARCHAR(10)" json:"co2_concentration"`
	Gettime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Humid       string    `xorm:"VARCHAR(10)" json:"humidity"`
	Illuminance string    `xorm:"VARCHAR(10)" json:"illuminance"`
	Machineid   string    `xorm:"not null VARCHAR(3)" json:"no"`
	No          int       `xorm:"not null pk autoincr INT(11)"`
	SoilHumid   string    `xorm:"VARCHAR(10)" json:"soil_humidity"`
	Temp        string    `xorm:"VARCHAR(10)" json:"temperture"`
	Wavelength  string    `xorm:"VARCHAR(10)" json:"wavelength"`
}
