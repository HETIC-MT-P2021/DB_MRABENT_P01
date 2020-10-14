package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"goevent/database"
	"goevent/models"
)



func GetOffice(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	offices, _ := repository.GetOfficeByCode(id)
	c.JSON(http.StatusOK, offices)
}
