package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/hatobus/ServerSmartAgri/PresenterDB"

	"gopkg.in/go-playground/validator.v9"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/hatobus/ServerSmartAgri/models"
)

var (
	ErrBadParameter = errors.New("Bad Parameter type")
	ErrInvalidJSON  = errors.New("Couldn't parse JSON")
)

type singleGetParameter struct {
	machine string `validate:"required"`
}

type multipleGetParameterWithDevice struct {
	machine string `validate:"required"`
	reqnum  int    `validate:"required, max=100,min=1"`
}

type multipleGetParameterWithTime struct {
	machine string `validate:"required"`
	reqnum  int    `validate:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/single", singlePOST)
	r.GET("/single/:machine", singleGET)
	r.POST("/multiple/device/:machine/:num", multiGETWithDevice)
	r.POST("/multiple/time/:machine", multiGETWithTime)
	r.Run()
}

func singlePOST(c *gin.Context) {

	iot := &models.Iotdata{}

	err := c.BindJSON(iot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrInvalidJSON,
			"Data":    iot,
		})
	}

	spew.Dump(iot)

	_, err = PresenterDB.StoreSingleIoTData(iot)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "InternalServerError",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Correctry stored",
	})
}

func singleGET(c *gin.Context) {
	v := validator.New()

	res := &models.Iotdata{}

	s := singleGetParameter{
		machine: c.Param("machine"),
	}

	if err := v.Struct(s); err != nil {
		log.Println("Illigal parameter")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrBadParameter,
			"Data":    res,
		})
		return
	}

	res, err := PresenterDB.RetSingleIoTData(s.machine)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "DBError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": res,
	})
}

func multiGETWithDevice(c *gin.Context) {
	v := validator.New()

	res := &[]models.Iotdata{}

	r, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrBadParameter,
		})
		return
	}

	m := multipleGetParameterWithDevice{
		machine: c.Param("machine"),
		reqnum:  r,
	}

	if err := v.Struct(m); err != nil {
		log.Println("Illigal parameter")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrBadParameter,
			"Data":    res,
		})
		return
	}

	res, err = PresenterDB.RetMultiIoTData(m.machine, m.reqnum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Your request is not correct.",
			"Data":    res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Found " + c.Param("num") + " records",
		"Data":    res,
	})
}

func multiGETWithTime(c *gin.Context) {
	v := validator.New()

	res := &[]models.Iotdata{}

	rt := &models.ReqestTimestamp{}

	err := c.BindJSON(rt)
	if err != nil {
		log.Println("Couldn't Bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrInvalidJSON,
			"Data":    res,
		})
		return
	}

	m := multipleGetParameterWithTime{
		machine: c.Param("machine"),
		reqnum:  rt.Limit,
	}

	if err := v.Struct(m); err != nil {
		log.Println("Illigal parameter")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": ErrBadParameter,
			"Data":    res,
		})
		return
	}

	res, err = PresenterDB.RetMultiIoTDataFromTime(m.machine, rt)
	if err != nil {
		log.Println("DB Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err,
			"Data":    res,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Found",
		"Data":    res,
	})
}
