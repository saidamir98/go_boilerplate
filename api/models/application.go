package models

import "time"

// Application ...
type Application struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateApplication ...
type CreateApplication struct {
	ID   string `json:"id" swaggerignore:"true"`
	Body string `json:"body"`
}

// ApplicationCreated ...
type ApplicationCreated struct {
	ID string `json:"id"`
}

// UpdateApplication ...
type UpdateApplication struct {
	ID   string `json:"id" swaggerignore:"true"`
	Body string `json:"body"`
}

// ApplicationList ...
type ApplicationList struct {
	Count        int           `json:"count"`
	Applications []Application `json:"applications"`
}

// ApplicationQueryParam ...
type ApplicationQueryParam struct {
	Search      string `json:"search"`
	Order       string `json:"order" enums:"id,body,created_at, updated_at"`
	Arrangement string `json:"arrangement" enums:"asc,desc"`
	Offset      int    `json:"offset" default:"0"`
	Limit       int    `json:"limit" default:"10"`
}
