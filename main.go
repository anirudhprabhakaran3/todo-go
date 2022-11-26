package main

import (
	"example/todo/controllers"
	"example/todo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/todos", controllers.FindTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.FindTodo)
	router.PATCH("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
