package entity

import "time"

// Booking struct merepresentasikan data booking tiket dalam database
type Booking struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	TicketID   string    `json:"ticket_id"`
	Amount     int       `json:"amount"` // Jumlah tiket
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"` // pending, confirmed, cancelled
	CreatedAt  time.Time `json:"created_at"`
}

// BookingRequest struct untuk validasi input dari user
type BookingRequest struct {
	UserID   string `json:"user_id"`
	TicketID string `json:"ticket_id"`
	Amount   int    `json:"amount"`
}
