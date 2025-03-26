package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Ajoute ici tous les champs nécessaires
}
