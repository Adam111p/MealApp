package main

import (
	"meal-order-app/backend/config"
	"meal-order-app/backend/database"
	"meal-order-app/backend/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	conf := config.LoadConf()

	database.InitDB(conf)
	// Grupowanie tras sprawia, że kod jest czytelny
	api := r.Group("/api")
	{
		api.GET("/menu", handlers.GetMenu)
		api.GET("/spices", handlers.GetSpices)
		api.GET("/searchDishes", handlers.SearchMenu)

		// api.POST("/order", handlers.CreateOrder) <-- tu dodasz kolejne
	}

	r.Run(":8080")
}
