package book

type Book struct {
	Id   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name" validate:"required"`
}
