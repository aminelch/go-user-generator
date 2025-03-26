package main

import (
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []User{
		{Uuid: "294e4f6c-5c7f-4956-a1d3-002f93d250e1", Name: "Amine Louhichi", Email: "q14wxka66@mozmail.com"},
		{Uuid: "a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8", Name: "Sophia Martinez", Email: "sophia.martinez@example.com"},
		{Uuid: "b2c3d4e5-f6g7-8901-h2i3-j4k5l6m7n8o9", Name: "Liam Johnson", Email: "liam.johnson@example.com"},
		{Uuid: "c3d4e5f6-g7h8-9012-i3j4-k5l6m7n8o9p0", Name: "Olivia Brown", Email: "olivia.brown@example.com"},
		{Uuid: "d4e5f6g7-h8i9-0123-j4k5-l6m7n8o9p0q1", Name: "Noah Davis", Email: "noah.davis@example.com"},
		{Uuid: "e5f6g7h8-i9j0-1234-k5l6-m7n8o9p0q1r2", Name: "Emma Wilson", Email: "emma.wilson@example.com"},
		{Uuid: "f6g7h8i9-j0k1-2345-l6m7-n8o9p0q1r2s3", Name: "Ava Garcia", Email: "ava.garcia@example.com"},
		{Uuid: "g7h8i9j0-k1l2-3456-m7n8-o9p0q1r2s3t4", Name: "William Miller", Email: "william.miller@example.com"},
		{Uuid: "h8i9j0k1-l2m3-4567-n8o9-p0q1r2s3t4u5", Name: "Isabella Lee", Email: "isabella.lee@example.com"},
		{Uuid: "i9j0k1l2-m3n4-5678-o9p0-q1r2s3t4u5v6", Name: "James Anderson", Email: "james.anderson@example.com"},
		{Uuid: "j0k1l2m3-n4o5-6789-p0q1-r2s3t4u5v6w7", Name: "Mia Taylor", Email: "mia.taylor@example.com"},
		{Uuid: "k1l2m3n4-o5p6-7890-q1r2-s3t4u5v6w7x8", Name: "Benjamin Thomas", Email: "benjamin.thomas@example.com"},
		{Uuid: "l2m3n4o5-p6q7-8901-r2s3-t4u5v6w7x8y9", Name: "Charlotte Moore", Email: "charlotte.moore@example.com"},
		{Uuid: "m3n4o5p6-q7r8-9012-s3t4-u5v6w7x8y9z0", Name: "Lucas Jackson", Email: "lucas.jackson@example.com"},
		{Uuid: "n4o5p6q7-r8s9-0123-t4u5-v6w7x8y9z0a1", Name: "Amelia White", Email: "amelia.white@example.com"},
		{Uuid: "o5p6q7r8-s9t0-1234-u5v6-w7x8y9z0a1b2", Name: "Henry Harris", Email: "henry.harris@example.com"},
		{Uuid: "p6q7r8s9-t0u1-2345-v6w7-x8y9z0a1b2c3", Name: "Evelyn Clark", Email: "evelyn.clark@example.com"},
		{Uuid: "q7r8s9t0-u1v2-3456-w7x8-y9z0a1b2c3d4", Name: "Alexander Lewis", Email: "alexander.lewis@example.com"},
		{Uuid: "r8s9t0u1-v2w3-4567-x8y9-z0a1b2c3d4e5", Name: "Harper Robinson", Email: "harper.robinson@example.com"},
		{Uuid: "s9t0u1v2-w3x4-5678-y9z0-a1b2c3d4e5f6", Name: "Michael Walker", Email: "michael.walker@example.com"},
		{Uuid: "t0u1v2w3-x4y5-6789-z0a1-b2c3d4e5f6g7", Name: "Abigail Young", Email: "abigail.young@example.com"},
	}
	if result := db.Create(&users); result.Error != nil {
		Logger.Fatal("Failed to seed users:", result.Error)
	}
	Logger.Info("Successfully seeded users")
}
