package model

import "http101/internal/domain/base"

type TaskModel struct {
	base.BaseModel `bson:",inline"`
	Name           string `bson:"name"`
	Description    string `bson:"description"`
}
