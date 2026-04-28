package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
	TypeBase struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"baseType"`
}
type BaseType int

const (
	SqlLite BaseType = iota
	PostgreSql
)

func LoadConf() Config {
	viper.SetConfigName("ConfigApp") // nazwa pliku bez rozszerzenia
	viper.SetConfigType("yaml")      // lub "json"
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Błąd odczytu pliku: %w", err))
	}

	// 2. Mapowanie na strukturę
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("Błąd mapowania: %w", err))
	}
	return conf

}

func GetDsnFromConfig(conf Config) string {
	// 3. Budowanie DSN dla GORM
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Dbname,
		conf.Database.Port,
	)
	return dsn
}

func ParseBaseType(s string) BaseType {
	switch s {
	case "postgres", "postgresql":
		return PostgreSql
	case "sqlite", "sqlite3":
		return SqlLite
	default:
		return SqlLite // domyślnie lub obsługa błędu
	}
}
