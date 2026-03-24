package models

import "time"

type Role string

const (
	RoleAdministrator Role = "administrator"
	RoleGuest         Role = "guest"
	RoleParent        Role = "parent"
	RoleChild         Role = "child"
	RoleDirector      Role = "director"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	FirstName string    `json:"first_name" gorm:"size:100;not null"`
	LastName  string    `json:"last_name" gorm:"size:100;not null"`
	Role      Role      `json:"role" gorm:"type:varchar(32);not null;default:guest"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
