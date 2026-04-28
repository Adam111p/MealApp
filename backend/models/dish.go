package models

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID           int             `json:"id" gorm:"primary_key;autoIncrement"`
	Name         string          `gorm:"unique;not null" json:"name"`
	Price        float64         `json:"price"`
	Description  string          `json:"description"`
	FileName     string          `json:"fileName"`
	TypeDish     string          `json:"typeDish"`
	Type         DishType        `gorm:"foreignKey:TypeDish;references:Name"`
	DishToppings []DishTopping   `json:"toppings"`
	Embedding    pgvector.Vector `gorm:"type:vector(768)"`
}
type DishTopping struct {
	Topping   Topping `json:"topping" gorm:"foreignKey:ToppingID"`
	DishID    int     `gorm:"primaryKey"`
	ToppingID int     `gorm:"primaryKey"`
	Quantity  int     `json:"quantity"`
}

type Topping struct {
	gorm.Model
	ID    int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name" gorm:"unique;not null"` // np. "Ser", "Cebula"
	Price float64 `json:"price"`                       // Cena za dodatek
}
type Spices struct {
	gorm.Model
	ID          int      `json:"id" gorm:"primary_key;autoIncrement"`
	Name        string   `gorm:not null" json:"name"`
	LevelSpice  int      `json:"levelSpice"`
	Description string   `json:"description"`
	TypeDish    string   `json:"typeDish"`
	Type        DishType `gorm:"foreignKey:TypeDish;references:Name"`
}

type DishType struct {
	// Ustawiamy Name jako Primary Key, żeby inne tabele mogły się do niego odnosić
	Name string `gorm:"primaryKey"`
}
