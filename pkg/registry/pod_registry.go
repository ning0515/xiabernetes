package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/client"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	. "github.com/learnk8s/xiabernetes/pkg/types"
	"net/url"
)

type PodRegistry struct {
	storage   PodStorage
	scheduler scheduler.Scheduler
}

func MakePodRegistry(storage PodStorage, scheduler scheduler.Scheduler) *PodRegistry {
	return &PodRegistry{
		storage:   storage,
		scheduler: scheduler,
	}
}

func (t *PodRegistry) Create(pod interface{}) {
	newPod := pod.(Pod)
	t.storage.CreatePod(newPod, t.scheduler.Schedule(newPod))
}

func (t *PodRegistry) List(url *url.URL) interface{} {
	var result PodList
	var query *map[string]string
	if url != nil {
		queryMap := client.StringToLabel(url.Query().Get("labels"))
		query = &queryMap
	}
	result = PodList{
		Items: t.storage.ListPod(query),
	}
	result.Kind = "cluster#podList"
	return result
}

func (t *PodRegistry) Extract(data []byte) interface{} {
	pod := Pod{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &pod)
	fmt.Printf("in Extract:\n %v\n", pod)
	return pod
}
