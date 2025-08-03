package handlers

import "gorm.io/gorm"

// APIHandler provides shared database access and common functionality for all API handlers.
type APIHandler struct {
	db *gorm.DB
}

func NewAPIHandler(db *gorm.DB) *APIHandler {
	return &APIHandler{db: db}
}
