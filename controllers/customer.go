package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"goevent/database"
	"goevent/models"

)


func AllCustomers(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetAllCustomers()
	c.JSON(http.StatusOK, customers)
}

func CustomerNum(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	customer, _ := repository.GetCustomerByNum(id)
	c.JSON(http.StatusOK, customer)
}
