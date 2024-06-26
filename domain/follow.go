package domain

import (
	"context"
	"time"
)

type Follow struct {
	FollowID   string `json:"_id" bson:"_id"`
	FollowerID string `json:"follower_id" bson:"follower_id"`
	FollowedID string `json:"followed_id" bson:"followed_id"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`

}

type FollowRepository interface {
	Create(c context.Context, follow *Follow) (*Follow, error)
	GetByFollowerID(c context.Context, followerID string) (*[]*Follow, error)
	GetByFollowedID(c context.Context, followedID string) (*[]*Follow, error)
	Delete(c context.Context, follow *Follow) (*Follow, error)
}