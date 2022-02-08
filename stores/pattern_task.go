package stores

import (
	"mathegeapi/models"
)

type PatternTaskStore struct {
	store *Store
}

func NewPatternTaskStore(store *Store) *PatternTaskStore {
	return &PatternTaskStore{store: store}
}

func (ts *PatternTaskStore) Get(id uint) (*models.PatternTask, error) {
	patternTask := models.PatternTask{}

	err := ts.store.Get(id, &patternTask)
	if err != nil {
		return nil, err
	}

	return &patternTask, nil
}

func (ts *PatternTaskStore) GetAll() (patternTasks models.PatternTasks, err error) {
	err = ts.store.GetAll(&patternTasks)
	if err != nil {
		return nil, err
	}

	return patternTasks, nil
}

func (ts *PatternTaskStore) Add(patternTask models.PatternTask) (*models.PatternTask, error) {
	err := ts.store.Add(&patternTask)
	if err != nil {
		return nil, err
	}

	return &patternTask, nil
}

func (ts *PatternTaskStore) Delete(id uint) {
	ts.store.Delete(&models.PatternTask{}, id)
}

func (ts *PatternTaskStore) Update(id uint, patternTask models.PatternTask) (*models.PatternTask, error) {
	err := ts.store.Update(id, &patternTask)
	if err != nil {
		return nil, err
	}

	return &patternTask, nil
}
