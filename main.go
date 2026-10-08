// ItemService
package main

import (
	"fmt"
	"itemservice/handler"
	"itemservice/model"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database configuration
var (
	dbHost     = "postgresdb"
	dbUser     = "bismark"
	dbPassword = "l%wTQkfWv?2_"
	dbName     = "itemservice"
)

func main() {
	// Connection string format (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&model.Item{}, &model.ReservationItem{})

	inventoryHandler := handler.CreateInventoryHandler(db)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/inventory/reserve", inventoryHandler.InventoryReserveHandler)
	v1.POST("/inventory/confirm", inventoryHandler.InventoryConfirmHandler)
	v1.GET("/inventory/stock", inventoryHandler.InventoryStatusHandler)

	router.Run(":8080")
}
