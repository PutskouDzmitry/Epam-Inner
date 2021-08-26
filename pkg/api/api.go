package api

import (
	"github.com/PutskouDzmitry/Epam-Inner/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bookAPI struct {
	data *data.EntityData
}

func ServeUserResource(data data.EntityData, router *gin.Engine) {
	api := &bookAPI{data: &data}

	router.GET("/get", api.getEntity)
}

func (b bookAPI) getEntity(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	space := c.Query("space")
	result := b.data.GetEntityDb(from, to, space)
	if result == -1 {
		logrus.Error("this data doesn't exist in db")
		c.JSON(500, gin.H{
			"message": "this data doesn't exist in db",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	}
}
