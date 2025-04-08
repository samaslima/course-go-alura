package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Name: "calabresa", Price: 40.5},
	{ID: 2, Name: "frango", Price: 35.4},
	{ID: 3, Name: "marguerita", Price: 30.4},
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func addPizzas(c *gin.Context) {
	var newPizza models.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pizzas = append(pizzas, newPizza)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}

	c.JSON(404, gin.H{"message": "pizza not found"})
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", addPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.Run()
}
