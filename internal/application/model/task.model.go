package model

import (
	base_model "http101/internal/application/model/base"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskModel struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Task string             `bson:"task" json:"task"`
	base_model.BaseModel
}
