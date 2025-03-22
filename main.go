package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
	time.Sleep(1000 * time.Millisecond)
	a, _ := strconv.Atoi(c.Param("a"))
	b, _ := strconv.Atoi(c.Param("b"))
	c.JSON(http.StatusOK, gin.H{
		"result": a + b,
	})
}

func sub(c *gin.Context) {
	time.Sleep(500 * time.Millisecond)
	a, _ := strconv.Atoi(c.Param("a"))
	b, _ := strconv.Atoi(c.Param("b"))
	c.JSON(http.StatusOK, gin.H{
		"result": a - b,
	})
}

type TwoNum struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func mul(c *gin.Context) {
	time.Sleep(700 * time.Millisecond)
	var two_num TwoNum

	if err := c.ShouldBindJSON(&two_num); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": two_num.A * two_num.B,
	})
}

func div(c *gin.Context) {
	time.Sleep(200 * time.Millisecond)
	a, _ := strconv.Atoi(c.Query("a"))
	b, _ := strconv.Atoi(c.Query("b"))
	c.JSON(http.StatusOK, gin.H{
		"result": a / b,
	})
}

func main() {
	r := gin.Default()
	r.GET("/add/:a/:b", add)
	r.GET("/sub/:a/:b", sub)
	r.POST("/mul", mul)
	r.GET("/div", div)
	r.Run(":8080")
}
