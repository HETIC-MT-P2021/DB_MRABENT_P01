package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"goevent/database"
	"goevent/models"
)

func GetAllEmployees(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	employees, _ := repository.GetEmployees()
	c.JSON(http.StatusOK, employees)
}

func GetByEmployee(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	employee, _ := repository.GetEmployeeByOffice(id)
	c.JSON(http.StatusOK, employee)
}


