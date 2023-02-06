package user

import (
	"time"
	"urlShortner/pkg/utils/hashing"

	"github.com/google/uuid"
)

type User struct {
	UserID    string `json:"userId" bson:"userId"`
	Username  string `json:"username" binding:"required" bson:"username"`
	Name      string `json:"name" binding:"required" bson:"name"`
	Email     string `json:"email" binding:"required,email" bson:"email"`
	Password  string `json:"password" binding:"required" bson:"password"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
	LastLogin int64  `json:"lastLogin" bson:"lastLogin"`
}

func (user *User) InitUser() User {
	user.Password = hashing.HashAndSalt(user.Password)
	user.UserID = uuid.NewString()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.LastLogin = time.Now().Unix()

	return *user
}

type Login struct {
	Email    string `json:"email" binding:"required,email" bson:"email"`
	Password string `json:"password" binding:"required" bson:"password"`
}
