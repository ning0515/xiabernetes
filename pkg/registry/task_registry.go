package registry

import (
	"encoding/json"
	"fmt"
	. "github.com/learnk8s/xiabernetes/pkg/types"
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
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &task)
	fmt.Printf("in Extract:\n %v\n", task)
	return task
}
