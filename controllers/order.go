package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goevent/database"
	"goevent/models"
)

func AllProduct(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	Details, _ := repository.GetAllProduct()
	c.JSON(http.StatusOK, Details)
}
