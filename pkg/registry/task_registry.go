package registry

import (
	"encoding/json"
	. "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

type TaskRegistry struct {
	storage TaskStorage
}

func MakeTaskRegistry(storage TaskStorage) *TaskRegistry {
	return &TaskRegistry{
		storage: storage,
	}
}

func (t *TaskRegistry) Create(name string) {
	t.storage.CreateTask(name)
}

func (t *TaskRegistry) Extract(data []byte) interface{} {
	task := Task{}
	json.Unmarshal(data, &task)
	return task
}
