package domain

import (
	"context"
	"time"
)


type User struct {
    UserID         string    `json:"user_id" bson:"_id"`
    Username       string    `json:"username" bson:"username"`
    Email          string    `json:"email" bson:"email"`
    Password       string    `json:"password" bson:"password"`
    Name           string    `json:"name" bson:"name"`
    Bio            string    `json:"bio" bson:"bio"`
    ProfilePicture string    `json:"profile_picture" bson:"profile_picture"`
    Role           string    `json:"role" bson:"role"`
    CreatedAt      time.Time `json:"timestamp" bson:"timestamp"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (*User, error)
	Delete(c context.Context, user *User) (*User, error)
	UpdateProfile(c context.Context, user *User) (*User, error)
	UpdateProfilePicture(c context.Context, user *User) (*User, error)
	UpdatePassword(c context.Context, user *User) (*User, error)
	UpdateUsername(c context.Context, user *User) (*User, error)
	UpdateEmail(c context.Context, user *User) (*User, error)
	GetAll(c context.Context, param string) (*[]*User, error)
	GetByUsername(c context.Context,username string) (*User, error)
	GetById(c context.Context, id string) (*User, error)
	GetRole(c context.Context, username string) (string, error)
	GetByEmail(c context.Context, email string) (*User, error)
}
