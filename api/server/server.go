package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"InnoTaxi/models"
)

func CheckRate(c *gin.Context) {
	var phone int32
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	userRate, err := models.UserCheckRate(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, userRate)
}

func CheckOrders(c *gin.Context) {

}
