package service

import (
	"backend-boking-ticket/internal/entity"
	"backend-boking-ticket/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

// BookingService mendefinisikan logic bisnis untuk booking
type BookingService interface {
	CreateBooking(req entity.BookingRequest) (entity.Booking, error)
	GetBookingByID(id string) (entity.Booking, error)
	GetAllBookings() ([]entity.Booking, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func (s *bookingService) CreateBooking(req entity.BookingRequest) (entity.Booking, error) {
	// Validasi input sederhana
	if req.Amount <= 0 {
		return entity.Booking{}, errors.New("amount must be greater than 0")
	}

	// Simulasi harga tiket (misal: harga 1 tiket = 50.000)
	pricePerTicket := 50000.0
	totalPrice := float64(req.Amount) * pricePerTicket

	// Buat object Booking
	booking := entity.Booking{
		ID:         uuid.New().String(),
		UserID:     req.UserID,
		TicketID:   req.TicketID,
		Amount:     req.Amount,
		TotalPrice: totalPrice,
		Status:     "confirmed", // Langsung confirmed untuk simulasi
		CreatedAt:  time.Now(),
	}

	// Simpan ke repository
	err := s.repo.Create(booking)
	if err != nil {
		return entity.Booking{}, err
	}

	return booking, nil
}

func (s *bookingService) GetBookingByID(id string) (entity.Booking, error) {
	return s.repo.FindByID(id)
}

func (s *bookingService) GetAllBookings() ([]entity.Booking, error) {
	return s.repo.FindAll()
}
