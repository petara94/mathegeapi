package stores

import (
	"errors"
	"fmt"
	"gorm.io/gorm/clause"
	"mathegeapi/models"
)

type TaskStore struct {
	store *Store
}

func NewTaskStore(store *Store) *TaskStore {
	return &TaskStore{store: store}
}

func (ts *TaskStore) Get(id uint) (*models.Task, error) {
	ts.store.RLock()
	defer ts.store.RUnlock()

	task := models.Task{}
	ts.store.DB.Preload(clause.Associations).Find(&task, id)
	if task.ID == 0 {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (ts *TaskStore) GetAll() (tasks models.Tasks) {
	ts.store.RLock()
	defer ts.store.RUnlock()

	ts.store.DB.Preload(clause.Associations).Find(&tasks)

	return tasks
}

func (ts *TaskStore) Add(task models.Task) (*models.Task, error) {
	ts.store.Lock()
	defer ts.store.Unlock()

	task.Model = models.Model{}

	res := ts.store.DB.Create(&task)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("задача не может быть добавлена: %v", res.Error))
	}

	return &task, nil
}

func (ts *TaskStore) Delete(id uint) {
	ts.store.Lock()
	defer ts.store.Unlock()

	ts.store.DB.Delete(&models.Task{}, id)
}

func (ts *TaskStore) Update(id uint, task models.Task) (*models.Task, error) {
	ts.store.Lock()
	defer ts.store.Unlock()

	task.ID = id

	res := ts.store.DB.Model(&task).Omit("id", "created_at", "deleted_id", "updated_at").Updates(task.Allowed())
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("задача не может быть обновлена: %v", res.Error))
	}

	ts.store.DB.Preload(clause.Associations).Find(&task, id)

	return &task, nil
}
