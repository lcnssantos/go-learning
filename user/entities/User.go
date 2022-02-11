package entities

import "time"

type User struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	EmailConfirmed bool      `json:"emailConfirmed"`
	Password       string    `json:"-"`
	IsActive       bool      `json:"isActive"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
