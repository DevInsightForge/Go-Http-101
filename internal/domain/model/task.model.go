package model

import "http101/internal/domain/base"

type TaskModel struct {
	base.BaseModel `bson:",inline"`
	Name           string `bson:"name"`
	Description    string `bson:"description"`
}

func (task *TaskModel) MapNewValues(dto TaskModel) *TaskModel {
	if dto.Name != "" {
		task.Name = dto.Name
	}
	if dto.Description != "" {
		task.Description = dto.Description
	}
	return task
}
