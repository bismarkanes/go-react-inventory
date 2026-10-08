package handler

import (
	"fmt"
	"itemservice/dto"
	"itemservice/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventoryHandler struct {
	Db *gorm.DB
}

func CreateInventoryHandler(db *gorm.DB) InventoryHandler {
	return InventoryHandler{
		Db: db,
	}
}

func getReservationExpiredDuration() time.Duration {
	return time.Minute * time.Duration(5)
}

func (ih *InventoryHandler) InventoryReserveHandler(c *gin.Context) {
	bodyRequest := dto.ReserveItemRequest{}

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	// 1. check if the item exist
	var items []model.Item
	result := ih.Db.Where("id = ?", bodyRequest.ItemID).Find(&items).Limit(1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  result.Error.Error(),
		})

		return
	}

	// 2. check if the reservation exist
	// 3. check if the reservation expired
	var reservations []model.ReservationItem
	result = ih.Db.
		Where("item_id = ?", bodyRequest.ItemID).
		Where("user_id", bodyRequest.UserID).
		Where("confirmed <> ?", true).
		Where("expires_at > ?", time.Now()).Find(&reservations).Limit(1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  result.Error.Error(),
		})

		return
	}

	// 4. If no reservation found
	if len(reservations) == 0 {
		// 5. Validate the item's stock
		var items []model.Item

		if result := ih.Db.
			Where("id = ?", bodyRequest.ItemID).
			Find(&items).Limit(1); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  result.Error.Error(),
			})
			return
		}

		item := items[0]
		if bodyRequest.Quantity > item.Stock {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  fmt.Sprintf("Item stock %d < reserve quantity %d", item.Stock, bodyRequest.Quantity),
			})
			return
		}

		// 6. Create the reservation item
		reserve := model.ReservationItem{
			UserID:    bodyRequest.UserID,
			ItemID:    strconv.Itoa(int(items[0].ID)),
			Quantity:  bodyRequest.Quantity,
			ExpiresAt: time.Now().Add(getReservationExpiredDuration()),
		}

		if result := ih.Db.Create(&reserve); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  result.Error.Error(),
			})

			return
		}
		reservations = append(reservations, reserve)
	}

	// 7. send the reservation response
	response := dto.ReserveItemResponse{}
	response.MapFromModel(reservations[0])

	c.JSON(http.StatusOK, response)
}

func (ih *InventoryHandler) InventoryConfirmHandler(c *gin.Context) {
	bodyRequest := dto.ReserveConfirmItemRequest{}

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	// 1. Check if the reservation is existed, not expired and not confirmed yet
	var reservations []model.ReservationItem
	result := ih.Db.
		Where("id = ?", bodyRequest.ReservationID).
		Where("confirmed <> ?", true).
		Where("expires_at > ?", time.Now()).Find(&reservations).Limit(1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  result.Error.Error(),
		})

		return
	}

	// 2. Create reservation confirm response
	response := dto.ReserveConfirmItemResponse{}

	// 3. Update reservation
	if len(reservations) > 0 {
		tx := ih.Db.Begin()

		reservation := reservations[0]
		reservation.Confirmed = true
		reservation.ConfirmAt = time.Now()

		result = tx.Updates(&reservation)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  result.Error.Error(),
			})

			return
		}

		// 4. Find the item record
		var items []model.Item
		result := ih.Db.
			Where("id = ?", reservation.ItemID).
			Find(&items).Limit(1)

		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  result.Error.Error(),
			})

			return
		}

		if len(items) == 0 {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "No item found",
			})

			return
		}

		// 5. Validate item stock against the request reserve quantity
		item := items[0]
		if item.Stock < reservation.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  fmt.Sprintf("Item stock %d < reserve quantity %d", item.Stock, reservation.Quantity),
			})

			return
		}

		// 6. Update item stock information
		item.ReservedStock = reservation.Quantity
		result = tx.Updates(&item)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  result.Error.Error(),
			})

			return
		}

		tx.Commit()

		response.MapFromModel(reservation)
	} else {
		// Found no reservation, give a proper response
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  "No reservation found",
		})

		return
	}

	c.JSON(http.StatusOK, response)
}

func (ih *InventoryHandler) InventoryStatusHandler(c *gin.Context) {
	itemId := c.Query("item_id")

	// 1. Find the item record
	var items []model.Item
	result := ih.Db.
		Where("id = ?", itemId).
		Find(&items).Limit(1)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  result.Error.Error(),
		})

		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  "No item found",
		})

		return
	}

	// 2. Create inventory status response
	response := dto.ItemInventoryStatusResponse{}
	response.MapFromModel(items[0])

	c.JSON(http.StatusOK, response)
}
