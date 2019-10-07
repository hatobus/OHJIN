package models

type Iotdata struct {
	No          int    `xorm:"not null pk autoincr INT(11)"`
	Machineid   string `xorm:"not null VARCHAR(3)" json:"no"`
	Temp        string `xorm:"VARCHAR(10)" json:"temperture"`
	Humid       string `xorm:"VARCHAR(10)" json:"humidity"`
	SoilHumid   string `xorm:"VARCHAR(10)" json:"soil_humidity"`
	Co2         string `xorm:"VARCHAR(10)" json:"co2_concentration"`
	Illuminance string `xorm:"VARCHAR(10)" json:"illuminance"`
	Wavelength  string `xorm:"VARCHAR(10)" json:"wavelength"`
}
