package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	tasks := []entity.Task{}
	query := r.db.Find(&tasks)
	if query.Error != nil {
		return nil, query.Error
	}
	return tasks, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	query := r.db.Create(task)
	if query.Error != nil {
		return 0, query.Error
	}
	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	query := r.db.First(&task, id)
	if query.Error != nil {
		return entity.Task{}, query.Error
	}
	return task, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var tasks []entity.Task
	query := r.db.Find(&tasks, "category_id = ?", catId)
	if query.Error != nil {
		return []entity.Task{}, query.Error
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	if task.CategoryID != 0 {
		query := r.db.Model(&task).Update("category_id", task.CategoryID)
		if query.Error != nil {
			return query.Error
		}
		return nil
	} else {
		query := r.db.Model(&task).Updates(entity.Task{Title: task.Title, Description: task.Description})
		if query.Error != nil {
			return query.Error
		}
		return nil
	}
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	query := r.db.Delete(&entity.Task{}, id)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
