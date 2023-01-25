package chapter

type Chapter struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty"`
	Title    string `json:"title" bson:"title" validate:"required"`
	Contents string `json:"contents" bson:"contents" validate:"required"`
	Book     string `json:"book" bson:"book" validate:"required"`
}
