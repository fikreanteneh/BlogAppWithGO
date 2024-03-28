package repository

import (
	"BlogApp/domain"
	"context"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	database   *mongo.Database
	collection string
}

func NewTagRepository(db *mongo.Database, collection string) domain.TagRepository {
	return &TagRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.TagRepository.
func (t *TagRepository) Create(c context.Context, tag *domain.Tag) (*domain.Tag, error) {
	tag.TagID = primitive.NewObjectID().Hex()
	_, err := t.database.Collection(t.collection).InsertOne(c, *tag)
	if err != nil{ 
	  return nil, err
	}
  
	return tag, nil
}

// Delete implements domain.TagRepository.
func (t *TagRepository) Delete(c context.Context, tagID string) (*domain.Tag, error) {

	filter := bson.M{"_id": tagID}
	var tag domain.Tag
	err := t.database.Collection(t.collection).FindOneAndDelete(c, filter).Decode(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// GetAll implements domain.TagRepository.
func (t *TagRepository) GetAll(c context.Context, param string) (*[]*domain.Tag, error) {
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: param , Options: "i"}}}

    // Perform the find operation
    cursor, err := t.database.Collection(t.collection).Find(c, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(c)

    // Iterate through the cursor and decode each document into a Tag struct
    var tags []*domain.Tag
    for cursor.Next(c) {
        var tag domain.Tag
        if err := cursor.Decode(&tag); err != nil {
			cursor.Close(c)
            return nil, err
        }
        tags = append(tags, &tag)
    }

    // Check if any error occurred during cursor iteration
    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return &tags, nil
}

// GetByID implements domain.TagRepository.
func (t *TagRepository) GetByID(c context.Context, tagID string) (*domain.Tag, error) {
	filter := bson.M{"_id": tagID}
	var tag domain.Tag
	err := t.database.Collection(t.collection).FindOne(c, filter).Decode(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil

}

// GetByName implements domain.TagRepository.
func (t *TagRepository) GetByName(c context.Context, name string) (*domain.Tag, error) {
	filter := bson.M{"name": name}
	var tag domain.Tag
	err := t.database.Collection(t.collection).FindOne(c, filter).Decode(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// Update implements domain.TagRepository.
func (t *TagRepository) Update(c context.Context, tag *domain.Tag) (*domain.Tag, error) {
	filter := bson.M{"_id": tag.TagID}
	update := bson.M{"$set": bson.M{"name": tag.Name}}
	_, err := t.database.Collection(t.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	return tag, nil

}


