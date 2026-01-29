package entity

import "time"

// User struct merepresentasikan data pengguna
type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"-"` // Password tidak dikembalikan di respon JSON
	CreatedAt time.Time `json:"created_at"`
}

// RegisterRequest struct input untuk pendaftaran
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest struct input untuk login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse format balasan saat sukses login/register
type AuthResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"` // Dipersiapkan untuk JWT nanti
}
