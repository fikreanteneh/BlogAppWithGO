package domain


type Tag  struct {
	TagID string `json:"_id" bson:"_id"`
	Name string  `json:"name" bson:"_id"`
}


type TagRepository interface {
	Create(tag Tag) (*Tag, error)
	GetAll(param string) (*[]*Tag, error)
	GetByID(tagID string) (*Tag, error)
	Update(tag Tag) (*Tag, error)
	Delete(tagID string) (*Tag, error)
	GetByName(name string) (*Tag, error)
}