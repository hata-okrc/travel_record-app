package repository

import (
	"fmt"

	"github.com/hata-okrc/travel_record-app.git/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITripRepository interface {
	GetAllTrips(trips *[]model.Trip, userId uint) error
	GetTripById(trip *model.Trip, userId uint, tripId uint) error
	CreateTrip(trip *model.Trip) error
	UpdateTrip(trip *model.Trip, userId uint, tripId uint) error
	DeleteTrip(userId uint, tripId uint) error
}

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) ITripRepository {
	return &tripRepository{db}
}

func (tr *tripRepository) GetAllTrips(trips *[]model.Trip, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(trips).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tripRepository) GetTripById(trip *model.Trip, userId uint, tripId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(trip, tripId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tripRepository) CreateTrip(trip *model.Trip) error {
	if err := tr.db.Create(trip).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tripRepository) UpdateTrip(trip *model.Trip, userId uint, tripId uint) error {
	result := tr.db.Model(trip).Clauses(clause.Returning{}).Where("id=? AND user_id=?", tripId, userId).Update("title", trip.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *tripRepository) DeleteTrip(userId uint, tripId uint) error {
	result := tr.db.Where("id=? AND user_id=?", tripId, userId).Delete(&model.Trip{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}