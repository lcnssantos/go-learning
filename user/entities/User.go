package entities

import "time"

type User struct {
	Id             string
	Name           string
	Email          string
	EmailConfirmed bool
	Password       string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
