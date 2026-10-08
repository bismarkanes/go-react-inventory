package dto

import (
	"itemservice/model"
	"strconv"
	"time"
)

type ReserveItemRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	ItemID   string `json:"item_id" binding:"required,numeric"`
	Quantity int    `json:"quantity" binding:"required"`
}

type ReserveItemResponse struct {
	Status        string    `json:"status"`
	ReservationID string    `json:"reservation_id"`
	ItemID        string    `json:"item_id"`
	UserID        string    `json:"user_id"`
	Quantity      int       `json:"quantity"`
	ExpiresAt     time.Time `json:"expires_at"`
}

func (rir *ReserveItemResponse) MapFromModel(reserveItem model.ReservationItem) {
	rir.Status = "success"
	rir.ReservationID = strconv.Itoa(int(reserveItem.ID))
	rir.ItemID = reserveItem.ItemID
	rir.UserID = reserveItem.UserID
	rir.Quantity = reserveItem.Quantity
	rir.ExpiresAt = reserveItem.ExpiresAt
}

type ReserveConfirmItemRequest struct {
	ReservationID string `json:"reservation_id" binding:"required,numeric"`
}

type ReserveConfirmItemResponse struct {
	Status        string    `json:"status"`
	ReservationID string    `json:"reservation_id"`
	ConfirmedAt   time.Time `json:"confirmed_at"`
}

func (rcir *ReserveConfirmItemResponse) MapFromModel(reserveItem model.ReservationItem) {
	rcir.Status = "success"
	rcir.ReservationID = strconv.Itoa(int(reserveItem.ID))
	rcir.ConfirmedAt = reserveItem.ConfirmAt
}

type ItemInventoryStatusResponse struct {
	ItemID         string `json:"item_id"`
	TotalStock     int    `json:"total_stock"`
	ReservedStock  int    `json:"reserved_stock"`
	AvailableStock int    `json:"available_stock"`
}

func (ii *ItemInventoryStatusResponse) MapFromModel(item model.Item) {
	ii.ItemID = strconv.Itoa(int(item.ID))
	ii.TotalStock = item.Stock
	ii.ReservedStock = item.ReservedStock
	ii.AvailableStock = item.Stock - item.ReservedStock
}
