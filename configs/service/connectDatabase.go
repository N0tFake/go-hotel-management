package service

import (
	"fmt"
	"log"

	client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	room "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Room"
	sale "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Sale"
	initializers "github.com/N0tFake/go-hotel-management/configs/initializations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *initializers.Config) {
	fmt.Println("> Connecting database...")
	fmt.Println(config.DBHost)
	fmt.Println(config.DBName)
	fmt.Println(config.DBPort)
	fmt.Println(config.DBPassword)
	fmt.Println(config.DBUser)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting database")
	}

	log.Println("> Connecting to the database")

	err = db.AutoMigrate(&room.Room{}, &client.Client{}, &sale.Sale{})
	if err != nil {
		panic("Error migrating")
	}

	DB = db
}
