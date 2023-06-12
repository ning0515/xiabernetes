package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/client"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	. "github.com/learnk8s/xiabernetes/pkg/types"
	"net/url"
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

func (t *TaskRegistry) List(url *url.URL) interface{} {
	var result TaskList
	var query *map[string]string
	if url != nil {
		queryMap := client.StringToLabel(url.Query().Get("labels"))
		query = &queryMap
	}
	result = TaskList{
		Items: t.storage.ListTask(query),
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
