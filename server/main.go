package main

import (
	"github.com/gin-gonic/gin"
	"github.com/todoapp/server02/controllers"
)

func main() {

	router := gin.Default()
	router.GET("/Alltask", controllers.Getalltasks)
	router.POST("/Create", controllers.Createtask)
	router.PUT("/Complete/:id", controllers.Complete)
	router.DELETE("/Delete/:id", controllers.Deleteone)

	router.Run(":9090")
}
