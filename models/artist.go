package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model        // using property gorm.Model automatically will be included: id, createdAt, updatedAt, deletedAt.
	Name       string `json:"name" validate:"required"`
	Album      string `json:"album"`
	Released   string `json:"released"`
	Country    string `json:"country"`
}

func ValidateArtist(artist *Artist) error {
	validate := validator.New()

	if err := validate.Struct(artist); err != nil {
		return err
	}
	return nil
}
