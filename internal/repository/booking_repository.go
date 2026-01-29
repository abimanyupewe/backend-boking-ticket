package repository

import (
	"backend-boking-ticket/internal/entity"

	"gorm.io/gorm"
)

// BookingRepository mendefinisikan kontrak untuk akses data booking
type BookingRepository interface {
	Create(booking entity.Booking) error
	FindByID(id string) (entity.Booking, error)
	FindAll() ([]entity.Booking, error)
}

// bookingRepository adalah implementasi database (PostgreSQL) menggunakan GORM
type bookingRepository struct {
	db *gorm.DB
}

// NewBookingRepository membuat instance repository baru dengan koneksi DB
func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (r *bookingRepository) Create(booking entity.Booking) error {
	return r.db.Create(&booking).Error
}

func (r *bookingRepository) FindByID(id string) (entity.Booking, error) {
	var booking entity.Booking
	err := r.db.Where("id = ?", id).First(&booking).Error
	if err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}

func (r *bookingRepository) FindAll() ([]entity.Booking, error) {
	var bookings []entity.Booking
	err := r.db.Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
