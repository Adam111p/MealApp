package database

import (
	"fmt"
	"math/rand"
	"meal-order-app/backend/config"
	"meal-order-app/backend/models"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/pgvector/pgvector-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.Config) {
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
	GenerateRandomVector(768)
	dishes := []models.Dish{
		{Name: "Pizza", Price: 35.0, Description: "Pyszna", FileName: "pizza.jpeg", TypeDish: "PIZZA",
			DishToppings: []models.DishTopping{
				{Topping: createdToppings["Ser Mozzarella"], Quantity: 1},
				{Topping: createdToppings["Sos Czosnkowy"], Quantity: 1},
			}, Embedding: GenerateRandomVector(768),
		},
		{Name: "Pasta", Price: 42.0, Description: "Włoska", FileName: "spagetti.jpeg", TypeDish: "PASTA",
			DishToppings: []models.DishTopping{}, Embedding: GenerateRandomVector(768),
		},
		{Name: "Pizza Margherita", Price: 35.00, Description: "Klasyk", FileName: "pizza.jpeg", TypeDish: "PIZZA", Embedding: GenerateRandomVector(768)},
		{Name: "Pasta Carbonara", Price: 42.50, Description: "Bez śmietany", FileName: "spagetti.jpeg", TypeDish: "PASTA", Embedding: GenerateRandomVector(768)},
		{Name: "Sałatka Cezar", Price: 29.00, Description: "Z kurczakiem",
			FileName: "cezar.jpeg", TypeDish: "SALATKA", DishToppings: []models.DishTopping{
				{Topping: createdToppings["Cebula"], Quantity: 1},
				{Topping: createdToppings["Boczek"]},
			}, Embedding: GenerateRandomVector(768)},
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
func GenerateRandomVector(dim int) pgvector.Vector {
	// Inicjalizacja ziarna (seed), aby za każdym razem były inne liczby
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	vec := make([]float32, dim)
	for i := range vec {
		// rand.Float32() zwraca liczbę od 0.0 do 1.0
		// Jeśli chcesz zakres od -1.0 do 1.0, użyj: r.Float32()*2 - 1
		vec[i] = r.Float32()
	}

	return pgvector.NewVector(vec)
}
