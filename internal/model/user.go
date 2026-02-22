package model

import (
	"time"

	"gorm.io/datatypes"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleAuthor UserRole = "author"
	RoleGuest  UserRole = "guest"
)

type User struct {
	ID           int            `gorm:"column:user_id;primaryKey"`
	Username     string         `gorm:"column:username;type:varchar(255);not null;unique"`
	DisplayName  string         `gorm:"column:display_name;type:varchar(255);not null"`
	AvatarURL    string         `gorm:"column:avatar_url;type:varchar(255);not null"`
	Bio          string         `gorm:"column:bio;type:text"`
	Role         UserRole       `gorm:"column:role;type:varchar(255);not null"`
	Email        string         `gorm:"column:email;type:varchar(255);not null;unique"`
	PasswordHash string         `gorm:"column:password_hash;type:varchar(255);not null"`
	SocialLinks  datatypes.JSON `gorm:"column:social_links;type:json"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
