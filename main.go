package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hatobus/ServerSmartAgri/PresenterDB"

	"gopkg.in/go-playground/validator.v9"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/hatobus/ServerSmartAgri/models"
)

type singleGetParameter struct {
	machine string `validate:"required"`
}

type multipleGetParameter struct {
	machine string `validate:"required"`
	reqnum  int    `validate:"required, max=100,min=1"`
}

func main() {
	r := gin.Default()
	r.POST("/single", singlePOST)
	r.GET("/single/:machine", singleGET)
	r.POST("/multiple/:machine/:num", multiGET)
	r.Run()
}

func singlePOST(c *gin.Context) {

	iot := &models.Iotdata{}

	err := c.BindJSON(iot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
	}

	iot.Gettime = time.Now()

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
			"Message": "Bad parameter type",
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

func multiGET(c *gin.Context) {
	v := validator.New()

	res := &[]models.Iotdata{}

	r, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Illigal parameter",
		})
		return
	}

	m := multipleGetParameter{
		machine: c.Param("machine"),
		reqnum:  r,
	}

	if err := v.Struct(m); err != nil {
		log.Println("Illigal parameter")
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad parameter type",
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
		"Message": "Find " + c.Param("num") + " records",
		"Data":    res,
	})
}
