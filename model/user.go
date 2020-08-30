package model

import (
	"time"
)

// User ...
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// APIError example
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// RevValueBase example
type RevValueBase struct {
	Status bool `json:"Status"`

	Err int32 `json:"Err"`
}

// RevValue example
type RevValue struct {
	RevValueBase

	Data int `json:"Data"`
}
