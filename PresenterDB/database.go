package PresenterDB

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/go-xorm/xorm"
	"github.com/hatobus/ServerSmartAgri/models"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupEngine() (*xorm.Engine, error) {
	envLoad()

	engine, err := xorm.NewEngine(os.Getenv("dbName"), os.Getenv("dataSourceName"))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return engine, nil
}

func StoreSingleIoTData(d *models.Iotdata) (int64, error) {

	engine, err := setupEngine()
	spew.Dump(err)

	affected, err := engine.Insert(d)
	if err != nil {
		return affected, err
	}

	return affected, nil
}

func RetSingleIoTData(ID string) (*models.Iotdata, error) {
	envLoad()

	engine, err := setupEngine()

	ret := &models.Iotdata{}

	_, err = engine.Where("machineid = ?", ID).Desc("no").Get(ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func RetMultiIoTData(ID string, num int) (*[]models.Iotdata, error) {
	envLoad()

	ret := &[]models.Iotdata{}

	engine, err := setupEngine()
	if err != nil {
		return ret, err
	}

	err = engine.Where("machineid = ?", ID).Desc("no").Limit(num, 0).Find(ret)
	if err != nil {
		return ret, err
	}

	return ret, err
}

func RetMultiIoTDataFromTime(ID string, d *models.ReqestTimestamp) (*[]models.Iotdata, error) {
	envLoad()

	ret := &[]models.Iotdata{}

	engine, err := setupEngine()
	if err != nil {
		return ret, err
	}

	err = engine.Where("machineis = ?", ID).
		And("gettime >= ?", d.Start).
		And("gettime <= ?", d.End).
		Desc("no").
		Limit(d.Limit).
		Find(ret)

	if err != nil {
		return ret, err
	}

	return ret, nil
}
