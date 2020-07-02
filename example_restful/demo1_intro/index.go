package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func GetUsername(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

func Post(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"result": u})
}

func PutPathParameter(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

func DeleteQuery(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func main() {
	r := gin.Default()

	r.GET("/", GetHelloWorld)
	r.GET("/username", GetUsername)
	r.POST("/", Post)
	r.PUT("/user/:username/:password", PutPathParameter)
	r.DELETE("/username", DeleteQuery)
	// Run the server
	r.Run()
}
