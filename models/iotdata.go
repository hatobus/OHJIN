package models

type Iotdata struct {
	Machineid   string `json:"no"`
	Temp        string `json:"temperture"`
	Humid       string `json:"humidity"`
	SoilHumid   string `json:"soil_humidity"`
	Co2         string `json:"co2_concentration"`
	Illuminance string `json:"illuminance"`
	Wavelength  string `json:"wavelength"`
}
