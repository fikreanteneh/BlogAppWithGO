package domain

type Follow struct {
	FollowID string `json:"_id" bson:"_id"`
	FollowerID string `json:"follower_id" bson:"follower_id"`
	FollowedID string `json:"followed_id" bson:"followed_id"`
}


type FollowRepository interface {
	Create(follow Follow) (*Follow, error)
	GetByFollowerID(followerID string) (*[]*Follow, error)
	GetByFollowedID(followedID string) (*[]*Follow, error)
	Delete(followID string) (*Follow, error)
}