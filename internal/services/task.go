package services

import (
	"sync"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"

	"github.com/TheTNB/panel/app/jobs"
	"github.com/TheTNB/panel/app/models"
)

var taskMap sync.Map

type TaskImpl struct {
}

func NewTaskImpl() *TaskImpl {
	return &TaskImpl{}
}

func (r *TaskImpl) Process(taskID uint) error {
	if err := r.markAsRunning(taskID); err != nil {
		return err
	}
	return facades.Queue().Job(&jobs.ProcessTask{}, []any{taskID}).Dispatch()
}

func (r *TaskImpl) DispatchWaiting() error {
	var tasks []models.Task
	if err := facades.Orm().Query().Where("status = ?", models.TaskStatusWaiting).Find(&tasks); err != nil {
		return err
	}

	for _, task := range tasks {
		if _, ok := taskMap.Load(task.ID); ok {
			continue
		}
		if err := r.Process(task.ID); err != nil {
			return err
		}
	}

	return nil
}

func (r *TaskImpl) markAsRunning(taskID uint) error {
	task := models.Task{
		Model:  orm.Model{ID: taskID},
		Status: models.TaskStatusRunning,
	}
	if _, err := facades.Orm().Query().Where("id", taskID).Update(&task); err != nil {
		return err
	}

	taskMap.Store(taskID, true)
	return nil
}
