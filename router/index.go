package router

import (
	"github.com/gin-gonic/gin"

	"goevent/controllers"
)

func addIndexRoutes(rg *gin.RouterGroup) {
	index := rg.Group("/")
	customers := rg.Group("/customers")
	orders := rg.Group("/orders")
	employees := rg.Group("/employees")
	offices := rg.Group("/offices")
	//ordersdetails := rg.Group("/ordersdetails")

	// Relative path
	index.GET("/", controllers.IndexAction)

	// Customers
	customers.GET("/", controllers.AllCustomers)
	customers.GET("/:id", controllers.CustomerNum)

	// Orders
	orders.GET("/", controllers.AllProduct)

	//Employees
	employees.GET("/", controllers.GetAllEmployees)
	employees.GET("/:id", controllers.GetByEmployee)

	//Offices
	offices.GET("/:id", controllers.GetOffice)


	//Order details

}