package handlers

import "gorm.io/gorm"

// BaseHandler provides shared database access and common functionality for all API handlers.
type BaseHandler struct {
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{db: db}
}
