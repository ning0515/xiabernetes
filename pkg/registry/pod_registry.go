package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	. "github.com/learnk8s/xiabernetes/pkg/types"
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

func (t *PodRegistry) Create(pod interface{}) <-chan interface{} {
	newPod := pod.(Pod)
	return apiserver.MakeAsync(func() interface{} {
		//time.Sleep(10 * time.Second)
		t.storage.CreatePod(newPod, t.scheduler.Schedule(newPod))
		fmt.Println("创建完成")
		return newPod
	})
}

func (t *PodRegistry) List(query labels.Query) interface{} {
	var result PodList
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
