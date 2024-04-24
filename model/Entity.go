package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
}
