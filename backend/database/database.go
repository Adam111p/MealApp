package database

import (
	"meal-order-app/backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open(sqlite.Open("LocalTesBase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Dish{}, &models.Spices{}, &models.DishType{})

	DB = db
	Seed()
}

// Test init
func Seed() {
	types := []models.DishType{
		{Name: "PIZZA"},
		{Name: "PASTA"},
		{Name: "SALATKA"},
	}
	for _, t := range types {
		DB.FirstOrCreate(&t, models.DishType{Name: t.Name})
	}

	dishes := []models.Dish{
		{Name: "Pizza", Price: 35.0, Description: "Pyszna", FileName: "pizza.jpeg", TypeDish: "PIZZA"},
		{Name: "Pasta", Price: 42.0, Description: "Włoska", FileName: "spagetti.jpeg", TypeDish: "PASTA"},
		{Name: "Pizza Margherita", Price: 35.00, Description: "Klasyk", FileName: "pizza.jpeg", TypeDish: "PIZZA"},
		{Name: "Pasta Carbonara", Price: 42.50, Description: "Bez śmietany", FileName: "spagetti.jpeg", TypeDish: "PASTA"},
		{Name: "Sałatka Cezar", Price: 29.00, Description: "Z kurczakiem", FileName: "cezar.jpeg", TypeDish: "SALATKA"},
	}
	for _, d := range dishes {
		DB.FirstOrCreate(&d, models.Dish{Name: d.Name})
	}

	spices := []models.Spices{
		{Name: "Papryka", LevelSpice: 4, Description: "papryka chilli", TypeDish: "PIZZA"},
		{Name: "Papryka", LevelSpice: 4, Description: "papryka chilli", TypeDish: "PASTA"},
		{Name: "Papryka", LevelSpice: 4, Description: "papryka chilli", TypeDish: "SALATKA"},

		{Name: "Czosnek", LevelSpice: 3, Description: "polski czosnek", TypeDish: "PIZZA"},
		{Name: "Czosnek", LevelSpice: 3, Description: "polski czosnek", TypeDish: "PASTA"},

		// Koperkowy tylko do Sałatek i Pizzy (jako sos)
		{Name: "koperkowy", LevelSpice: 2, Description: "Koperek", TypeDish: "SALATKA"},
		{Name: "koperkowy", LevelSpice: 2, Description: "Koperek", TypeDish: "PIZZA"},
	}

	for _, s := range spices {
		DB.FirstOrCreate(&s, models.Spices{Name: s.Name, TypeDish: s.TypeDish})
	}
}
