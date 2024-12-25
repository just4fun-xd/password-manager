package models

import "time"

type Password struct {
	ID          int       `json:"id"`
	ServiceName string    `json:"service_name"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}
