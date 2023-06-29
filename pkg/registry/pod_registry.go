package registry

import (
	"encoding/json"
	"fmt"
	. "github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
)

type PodRegistryStorage struct {
	storage   PodRegistry
	scheduler scheduler.Scheduler
}

func MakePodRegistryStorage(storage PodRegistry, scheduler scheduler.Scheduler) *PodRegistryStorage {
	return &PodRegistryStorage{
		storage:   storage,
		scheduler: scheduler,
	}
}

func (t *PodRegistryStorage) Create(pod interface{}) <-chan interface{} {
	newPod := pod.(Pod)
	return apiserver.MakeAsync(func() interface{} {
		//time.Sleep(10 * time.Second)
		t.storage.CreatePod(newPod, t.scheduler.Schedule(newPod))
		fmt.Println("创建完成")
		return newPod
	})
}

func (t *PodRegistryStorage) List(query labels.Query) interface{} {
	var result PodList
	result = PodList{
		Items: t.storage.ListPod(query),
	}
	return result
}

func (t *PodRegistryStorage) Extract(data []byte) interface{} {
	pod := Pod{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &pod)
	fmt.Printf("in Extract:\n %v\n", pod)
	return pod
}
