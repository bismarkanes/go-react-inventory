// ItemService
package main

import (
	"itemservice/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InventoryReserveHandler(c *gin.Context) {
	bodyRequest := dto.ReserveItemRequest{}

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   bodyRequest,
	})
}

func InventoryConfirmHandler(c *gin.Context) {
	bodyRequest := dto.ReserveConfirmItemRequest{}

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   bodyRequest,
	})
}

func InventoryStatusHandler(c *gin.Context) {
	itemId := c.Query("item_id")

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"itemId": itemId,
	})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/inventory/reserve", InventoryReserveHandler)
	v1.POST("/inventory/confirm", InventoryConfirmHandler)
	v1.GET("/inventory/stock", InventoryStatusHandler)

	router.Run(":8080")
}
