package repository

import (
	"http101/internal/domain/model"
	base_repository "http101/internal/infrastructure/repository/base"
)

type TaskRepository struct {
	*base_repository.BaseRepository[model.TaskModel]
}

func NewTaskRepository() *TaskRepository {
	repo, err := base_repository.NewBaseRepository[model.TaskModel]("tasks")
	if err != nil {
		panic(err)
	}
	return &TaskRepository{repo}
}
