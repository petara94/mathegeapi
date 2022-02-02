package stores

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"mathegeapi/config"
	"mathegeapi/models"
	"testing"
	"time"
)

func StoreForTests() *Store {
	cnf, err := config.LoadConfig("../" + config.ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	store := NewStore(&cnf.Database)
	err = store.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = store.DB.AutoMigrate(&models.Task{}, &models.TaskImage{}, &models.PatternTask{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return store
}

func TestTaskStore_Get(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := StoreForTests()
	task := models.Task{
		MathEgeID: uint(rand.Int()),
		TaskText:  "1233333",
	}

	s.DB.Create(&task)

	ts := NewTaskStore(s)

	get, err := ts.Get(task.ID)
	assert.Nil(t, err)
	assert.Equal(t, get.MathEgeID, task.MathEgeID)

	s.DB.Delete(&task)
}

func TestTaskStore_Add(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := StoreForTests()
	task := models.Task{
		MathEgeID: uint(rand.Int()),
		TaskText:  "1233333",
	}

	ts := NewTaskStore(s)
	get, err := ts.Add(task)
	assert.Nil(t, err)
	assert.NotEqual(t, get.ID, 0)

	s.DB.Unscoped().Delete(&get)
}

func TestTaskStore_Update(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := StoreForTests()
	task := models.Task{
		MathEgeID: uint(rand.Int()),
		TaskText:  "111",
	}

	ts := NewTaskStore(s)
	get, _ := ts.Add(task)

	updated, err := ts.Update(get.ID, models.Task{TaskText: "222"})
	assert.Nil(t, err)
	assert.Equal(t, updated.ID, get.ID)
	assert.Equal(t, updated.TaskText, "222")

	s.DB.Unscoped().Delete(&get)
}
