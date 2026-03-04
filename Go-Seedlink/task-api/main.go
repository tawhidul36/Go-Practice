package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{}
var idCounter = 1

func main() {
	r := gin.Default()

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

func createTask(c *gin.Context) {
	var task Task
	c.BindJSON(&task)
	task.ID = idCounter
	idCounter++
	task.Title = task.title
	tasks = append(tasks, task)
	c.JSON(http.StatusCreated, task)
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func updateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask Task
	c.BindJSON(&updatedTask)

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

func deleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
