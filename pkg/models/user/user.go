package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID   string `json:"userId" bson:"userId"`
	Username string `json:"username" binding:"required" bson:"username"`
	Name     string `json:"name" binding:"required" bson:"name"`
	Email    string `json:"email" binding:"required,email" bson:"email"`
	Password string `json:"password" binding:"required" bson:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (user *User) InitUser() User{
	user.UserID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return *user
}