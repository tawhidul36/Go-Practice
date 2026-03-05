package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task struct
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// in-memory storage
var tasks = []Task{}
var idCounter = 1

func main() {

	r := gin.Default()

	// routes
	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

// create new task
func createTask(c *gin.Context) {

	var task Task

	// bind JSON request
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// assign ID
	task.ID = idCounter
	idCounter++

	// save task
	tasks = append(tasks, task)

	c.JSON(http.StatusCreated, task)
}

// get all tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// update task
func updateTask(c *gin.Context) {

	// convert id param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTask Task

	// bind updated data
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// find and update
	for i, task := range tasks {
		if task.ID == id {

			tasks[i].Title = updatedTask.Title
			tasks[i].Done = updatedTask.Done

			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// delete task
func deleteTask(c *gin.Context) {

	// convert id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// find and delete
	for i, task := range tasks {
		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}