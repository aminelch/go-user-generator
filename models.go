package main

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Uuid  string `json:"uuid"`
	Email string `json:"email"`
}
