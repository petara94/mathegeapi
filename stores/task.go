package stores

import (
	"mathegeapi/models"
)

type TaskStore struct {
	store *Store
}

func NewTaskStore(store *Store) *TaskStore {
	return &TaskStore{store: store}
}

func (ts *TaskStore) Get(id uint) (*models.Task, error) {
	task := models.Task{}

	err := ts.store.Get(id, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (ts *TaskStore) GetAll() (tasks models.Tasks, err error) {
	err = ts.store.GetAll(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskStore) Add(task models.Task) (*models.Task, error) {
	err := ts.store.Add(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (ts *TaskStore) Delete(id uint) {
	ts.store.Delete(&models.Task{}, id)
}

func (ts *TaskStore) Update(id uint, task models.Task) (*models.Task, error) {
	err := ts.store.Update(id, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
