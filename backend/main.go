package main

import (
	"meal-order-app/backend/config"
	"meal-order-app/backend/database"
	"meal-order-app/backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conf := config.LoadConf()

	database.InitDB(conf)
	// Grupowanie tras sprawia, że kod jest czytelny
	api := r.Group("/api")
	{
		api.GET("/menu", handlers.GetMenu)
		api.GET("/spices", handlers.GetSpices)
		// api.POST("/order", handlers.CreateOrder) <-- tu dodasz kolejne
	}

	r.Run(":8080")
}
