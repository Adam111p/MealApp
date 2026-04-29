package handlers

import (
	"meal-order-app/backend/database"
	"meal-order-app/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMenu zwraca listę dostępnych dań
func GetMenu(c *gin.Context) {
	var dishes []models.Dish
	database.DB.Preload("DishToppings.Topping").Find(&dishes)

	c.JSON(http.StatusOK, dishes)
}

func GetSpices(c *gin.Context) {
	var spice []models.Spices
	database.DB.Find(&spice)

	c.JSON(http.StatusOK, spice)
}

func SearchMenu(c *gin.Context) {
	searchTerm := c.Query("query")
	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query is empty"})
		return
	}
	dishesSearch := database.SearchByDesc(searchTerm, c)
	c.JSON(http.StatusOK, dishesSearch)

}
