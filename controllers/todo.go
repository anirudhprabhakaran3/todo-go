package controllers

import (
	"example/todo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /todos
// Get all todo items

func FindTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// POST /todos
// Create new todo
func CreateTodo(c *gin.Context) {
	var input models.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Task: input.Task, Complete: input.Complete, CreatedAt: time.Now()}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GET /todos/:id
// Get todo of particular ID
func FindTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// PATCH /todos/:id
// Update a task as completed
func UpdateTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DELETE /todos/:id
// Delete a todo
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
