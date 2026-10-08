package dto

type ReserveItemRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	ItemID   string `json:"item_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type ReserveConfirmItemRequest struct {
	ReservationID string `json:"reservation_id" binding:"required"`
}
