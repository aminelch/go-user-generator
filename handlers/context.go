package handlers

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HandlerContext struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

var (
	hc *HandlerContext
)

func InitHandlers(logger *logrus.Logger, db *gorm.DB) {
	hc = &HandlerContext{
		Logger: logger,
		DB:     db,
	}
}
