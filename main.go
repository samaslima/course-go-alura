package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Name: "calabresa", Price: 40.5},
		{ID: 2, Name: "frango", Price: 35.4},
		{ID: 3, Name: "marguerita", Price: 30.4},
	}

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}
