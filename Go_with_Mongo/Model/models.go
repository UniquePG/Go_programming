package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie"`	// we can also write ",omitempty" if we want required field
	Director   string             `json:"director"`
	Year   int64             `json:"year"`
	Watched bool               `json:"watched"`
}
