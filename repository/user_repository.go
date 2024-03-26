package repository

import (
	"BlogApp/domain"
	"context"
	// "errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
  database   *mongo.Database
  collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
  return &UserRepository{
    database:   db,
    collection: collection,
  }
}

// Create implements domain.UserRepository.
func (u *UserRepository) Create(c context.Context, user *domain.User) (*domain.User, error) {
	user.UserID = primitive.NewObjectID().Hex()

	result, err := u.database.Collection(u.collection).InsertOne(c, user)
	if err != nil{ 
		return nil, err
	}

  newUser := &domain.User{
    UserID:           result.InsertedID.(primitive.ObjectID).Hex(),
    Email:            user.Email,
    Name:             user.Name,
    Bio:              user.Bio,
    Role:             user.Role,
    CreatedAt:        time.Now(),
  }


	return newUser, nil
}


// Delete implements domain.UserRepository.
func (u *UserRepository) Delete(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	_, err := u.database.Collection(u.collection).DeleteOne(c, filter)
	if err != nil {
        return nil, err
    }
	return user, nil
}

// GetAll implements domain.UserRepository.
func (u *UserRepository) GetAll(c context.Context, param string) (*[]*domain.User, error) {
    filter := bson.M{
        "$or": []bson.M{
            {"name": bson.M{"$regex": primitive.Regex{Pattern: param, Options: "i"}}},
            {"username": bson.M{"$regex": primitive.Regex{Pattern: param, Options: "i"}}},
        },
    }

    // Perform the find operation
    cursor, err := u.database.Collection(u.collection).Find(c, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(c)

    // Iterate through the cursor and decode each document into a User struct
    var users []*domain.User
    for cursor.Next(c) {
        var user domain.User
        if err := cursor.Decode(&user); err != nil {
			cursor.Close(c)
            return nil, err
        }
        users = append(users, &user)
    }

    // Check if any error occurred during cursor iteration
    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return &users, nil
}


// GetByEmail implements domain.UserRepository.
func (u *UserRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	result := u.database.Collection(u.collection).FindOne(c, filter)
	var user domain.User
	err := result.Decode(&user);
	if err != nil {
		return nil, err	
	}
	return &user, nil
}

// GetById implements domain.UserRepository.
func (u *UserRepository) GetById(c context.Context, id string) (*domain.User, error) {
	filter := bson.M{"_id": id}
	result := u.database.Collection(u.collection).FindOne(c, filter)		
	var user domain.User
	err := result.Decode(&user);
	if err != nil {
		return nil, err	
	}
	return &user, nil
}

// GetByUsername implements domain.UserRepository.
func (u *UserRepository) GetByUsername(c context.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	result := u.database.Collection(u.collection).FindOne(c, filter)
	var user domain.User
	err := result.Decode(&user);
	if err != nil {
		return nil, err	
	}
	return &user, nil
}

// GetRole implements domain.UserRepository.
func (u *UserRepository) GetRole(c context.Context, username string) (string, error) {
	filter := bson.M{"username": username}
	result := u.database.Collection(u.collection).FindOne(c, filter)
	var user domain.User
	err := result.Decode(&user);
	if err != nil {
		return "", err	
	}
	return user.Role, nil
}

// UpdateEmail implements domain.UserRepository.
func (u *UserRepository) UpdateEmail(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"email": user.Email}}
	_, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdatePassword implements domain.UserRepository.
func (u *UserRepository) UpdatePassword(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	_, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)	
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateProfile implements domain.UserRepository.
func (u *UserRepository) UpdateProfile(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"name": user.Name, "bio": user.Bio}}
	_, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateProfilePicture implements domain.UserRepository.
func (u *UserRepository) UpdateProfilePicture(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"profile_picture": user.ProfilePicture}}
	_, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUsername implements domain.UserRepository.
func (u *UserRepository) UpdateUsername(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"username": user.Username}}
	_, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}
