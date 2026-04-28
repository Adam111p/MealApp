package database

import (
	"fmt"
	"meal-order-app/backend/config"
	geminiApi "meal-order-app/backend/gemini"
	"meal-order-app/backend/models"

	"github.com/glebarez/sqlite"
	"github.com/pgvector/pgvector-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB
var Cfg config.Config

func InitDB(cfg config.Config) {
	Cfg = cfg
	var db *gorm.DB
	var err error

	dbType := config.ParseBaseType(cfg.TypeBase.Type)

	switch dbType {
	case config.SqlLite:

		db, err = gorm.Open(sqlite.Open("LocalTesBase.db"), &gorm.Config{})

	case config.PostgreSql:

		dsn := config.GetDsnFromConfig(cfg)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		fmt.Println("Połączenie z PostgreSQL udane! " + dsn)
	default:
		fmt.Println("Nieznany status")

	}

	if err != nil {
		panic("Nie udało się połączyć z bazą danych: " + err.Error())
	}
	DB = db

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error
	if err != nil {
		panic(fmt.Errorf("nie udało się włączyć pgvector: %v", err))
	}

	db.AutoMigrate(&models.Dish{}, &models.Spices{}, &models.DishType{}, &models.Topping{}, &models.DishTopping{})

	err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_dishes_embedding 
        ON dishes 
        USING hnsw (embedding vector_cosine_ops)`).Error
	if err != nil {
		panic(fmt.Errorf("nie udało się utworzyć indeksów hnsw: %v", err))
	}

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
	toppings := []models.Topping{
		{Name: "Ser Mozzarella", Price: 5.0},
		{Name: "Cebula", Price: 2.0},
		{Name: "Papryczki Jalapeño", Price: 3.5},
		{Name: "Boczek", Price: 6.0},
		{Name: "Sos Czosnkowy", Price: 2.5},
	}

	createdToppings := make(map[string]models.Topping)
	for _, top := range toppings {
		DB.FirstOrCreate(&top, models.Topping{Name: top.Name})
		createdToppings[top.Name] = top
	}

	dishes := []models.Dish{
		createDish("Pizza", 35.0, "Klasyczna włoska pizza na cienkim cieście z autorskim sosem pomidorowym, świeżą bazylią i nutą oliwy z oliwek", "pizza.jpeg", "PIZZA", []models.DishTopping{
			{Topping: createdToppings["Ser Mozzarella"], Quantity: 1},
			{Topping: createdToppings["Sos Czosnkowy"], Quantity: 1},
		}),
		createDish("Pasta", 42.0, "Tradycyjny makaron pszenny podawany z kremowym sosem, aromatycznym czosnkiem i kompozycją świeżo mielonych ziół", "spagetti.jpeg", "PASTA", []models.DishTopping{
			{Topping: createdToppings["Ser Mozzarella"], Quantity: 1},
			{Topping: createdToppings["Sos Czosnkowy"], Quantity: 1},
		}),
		createDish("Sałatka", 29.0, "Lekka i pożywna kompozycja chrupiących warzyw sezonowych, mixu sałat oraz aromatycznych dodatków z domowym vinegretem", "cezar.jpeg", "SALATKA", []models.DishTopping{
			{Topping: createdToppings["Cebula"], Quantity: 1},
			{Topping: createdToppings["Boczek"], Quantity: 1},
		}),
		createDish("Pizza Pepperoni", 40.0, "Pikantna uczta dla fanów wyrazistych smaków: plastry dojrzewającego pepperoni na grubym pokładzie ciągnącego sera", "pizza_pepperoni.jpeg", "PIZZA", []models.DishTopping{
			{Topping: createdToppings["Ser Mozzarella"], Quantity: 1},
			{Topping: createdToppings["Papryczki Jalapeño"], Quantity: 1},
		}),
		createDish("Pasta Bolognese", 45.0, "Sycące spaghetti z wolno gotowanym sosem mięsno-pomidorowym według oryginalnej receptury z Bolonii", "pasta_bolognese.jpeg", "PASTA", []models.DishTopping{
			{Topping: createdToppings["Ser Mozzarella"], Quantity: 1},
			{Topping: createdToppings["Sos Czosnkowy"], Quantity: 1},
		}),
		createDish("Sałatka Grecka", 27.0, "Świeże ogórki, dojrzałe pomidory, słona feta i czarne oliwki, skąpane w aromatycznym oregano i oliwie", "salatka_grecka.jpeg", "SALATKA", []models.DishTopping{
			{Topping: createdToppings["Cebula"], Quantity: 1},
			{Topping: createdToppings["Papryczki Jalapeño"], Quantity: 1},
		}),
		createDish("Pizza Margherita", 27.0, "Królowa prostoty: idealna harmonia czerwonych pomidorów, białej mozzarelli i zielonej, świeżej bazylii", "pizza.jpeg", "PIZZA", []models.DishTopping{}),
		createDish("Pasta Carbonara", 42.50, "Prawdziwie rzymski smak: sos na bazie żółtek jaj, twardego sera i chrupiącego boczku bez dodatku śmietany", "spagetti.jpeg", "PASTA", []models.DishTopping{}),
		createDish("Sałatka Cezar", 29.00, "Kultowa sałata rzymska z soczystym grillowanym kurczakiem, kruchymi grzankami i charakterystycznym sosem parmezanowym", "cezar.jpeg", "SALATKA", []models.DishTopping{
			{Topping: createdToppings["Cebula"], Quantity: 1},
			{Topping: createdToppings["Boczek"], Quantity: 1},
		}),
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

func createDish(name string, price float64, description string, fileName string, typeDish string, toppings []models.DishTopping) models.Dish {

	var values, error = geminiApi.GetEmbedding(Cfg.Gemini.ApiKey, description)
	if error != nil {
		panic(fmt.Errorf("Błąd pobierania embeddingu: %v", error))
	}

	dish := models.Dish{
		Name:         name,
		Price:        price,
		Description:  description,
		FileName:     fileName,
		TypeDish:     typeDish,
		DishToppings: toppings,
		Embedding:    pgvector.NewVector(values),
	}
	return dish
}
func SearchByDesc(description string) []models.Dish {
	var dishesSearch []models.Dish
	var queryVector, err = geminiApi.GetEmbedding(Cfg.Gemini.ApiKey, description)
	if err != nil {
		panic(fmt.Errorf("Błąd pobierania embeddingu: %v", err))
	}
	threshold := 0.4
	err = DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:  "embedding <=> ?",
			Vars: []interface{}{pgvector.NewVector(queryVector)},
		},
	}).
		Where("embedding <=> ? < ?", pgvector.NewVector(queryVector), threshold).
		Limit(5).Find(&dishesSearch).Error

	if err != nil {
		panic(fmt.Errorf("DB failed: %v", err))
	}
	return dishesSearch
}
