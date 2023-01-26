package books

type Book struct {
	Id   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name" validate:"required"`
}
