package main

import "gorm.io/gorm"

func SeedUsers(db *gorm.DB) {
    users := []User{
        {Name: "Amine Louhichi", Email: "q14wxka66@mozmail.com"},
        {Name: "Sophia Martinez", Email: "sophia.martinez@example.com"},
        {Name: "Liam Johnson", Email: "liam.johnson@example.com"},
        {Name: "Olivia Brown", Email: "olivia.brown@example.com"},
        {Name: "Noah Davis", Email: "noah.davis@example.com"},
        {Name: "Emma Wilson", Email: "emma.wilson@example.com"},
        {Name: "Ava Garcia", Email: "ava.garcia@example.com"},
        {Name: "William Miller", Email: "william.miller@example.com"},
        {Name: "Isabella Lee", Email: "isabella.lee@example.com"},
        {Name: "James Anderson", Email: "james.anderson@example.com"},
        {Name: "Mia Taylor", Email: "mia.taylor@example.com"},
        {Name: "Benjamin Thomas", Email: "benjamin.thomas@example.com"},
        {Name: "Charlotte Moore", Email: "charlotte.moore@example.com"},
        {Name: "Lucas Jackson", Email: "lucas.jackson@example.com"},
        {Name: "Amelia White", Email: "amelia.white@example.com"},
        {Name: "Henry Harris", Email: "henry.harris@example.com"},
        {Name: "Evelyn Clark", Email: "evelyn.clark@example.com"},
        {Name: "Alexander Lewis", Email: "alexander.lewis@example.com"},
        {Name: "Harper Robinson", Email: "harper.robinson@example.com"},
        {Name: "Michael Walker", Email: "michael.walker@example.com"},
        {Name: "Abigail Young", Email: "abigail.young@example.com"},
    }
    db.Create(&users)
}
