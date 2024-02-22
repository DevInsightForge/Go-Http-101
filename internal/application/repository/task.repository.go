package repository

import (
	"http101/internal/application/model"
	base_repository "http101/internal/application/repository/base"
)

type TaskRepository struct {
	*base_repository.BaseRepository[model.TaskModel]
}

func NewTaskRepository() (*TaskRepository, error) {
	repo, err := base_repository.NewBaseRepository[model.TaskModel]("tasks")
	if err != nil {
		return nil, err
	}
	return &TaskRepository{repo}, nil
}
