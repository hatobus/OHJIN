package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hatobus/ServerSmartAgri/PresenterDB"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/hatobus/ServerSmartAgri/models"
)

func main() {
	r := gin.Default()
	r.POST("/single", singlePOST)
	r.GET("/single/:machine", singleGET)
	r.Run()
}

func singlePOST(c *gin.Context) {

	iot := &models.Iotdata{}

	c.BindJSON(iot)
	iot.Gettime = time.Now()

	spew.Dump(iot)

	_, err := PresenterDB.StoreSingleIoTData(iot)
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
	res := &models.Iotdata{}

	machineNo := c.Param("machine")
	res, err := PresenterDB.RetSingleIoTData(machineNo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "DBError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": res,
	})
}
