package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	. "github.com/learnk8s/xiabernetes/pkg/types"
)

type TaskRegistry struct {
	storage   TaskStorage
	scheduler scheduler.Scheduler
}

func MakeTaskRegistry(storage TaskStorage, scheduler scheduler.Scheduler) *TaskRegistry {
	return &TaskRegistry{
		storage:   storage,
		scheduler: scheduler,
	}
}

func (t *TaskRegistry) Create(task interface{}) {
	newTask := task.(Task)
	t.storage.CreateTask(newTask, t.scheduler.Schedule(newTask))
}

func (t *TaskRegistry) List() interface{} {
	var result TaskList
	result = TaskList{
		Items: t.storage.ListTask(),
	}
	return result
}

func (t *TaskRegistry) Extract(data []byte) interface{} {
	task := Task{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &task)
	fmt.Printf("in Extract:\n %v\n", task)
	return task
}
