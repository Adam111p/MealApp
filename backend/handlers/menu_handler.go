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
	database.DB.Find(&dishes)

	c.JSON(http.StatusOK, dishes)
}

func GetSpices(c *gin.Context) {
	var spice []models.Spices
	database.DB.Find(&spice)

	c.JSON(http.StatusOK, spice)
}
