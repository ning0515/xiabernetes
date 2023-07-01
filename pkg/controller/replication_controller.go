package controller

import (
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/client"
	registry2 "github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	"math/rand"
)

type ReplicationManager struct {
	registry  registry2.WinRegistry
	scheduler scheduler.Scheduler
	client    client.ClientInterface
}

func MakeReplicateManager(registry registry2.WinRegistry, scheduler scheduler.Scheduler, client client.ClientInterface) *ReplicationManager {
	return &ReplicationManager{
		registry:  registry,
		scheduler: scheduler,
		client:    client,
	}
}
func (rm *ReplicationManager) Sync() {
	replicateControllers := rm.client.ListController()
	//fmt.Println(replicateControllers)
	for _, replicateController := range replicateControllers.Items {
		//fmt.Println("2223")
		rm.syncReplicationController(replicateController)
	}
}

func (rm *ReplicationManager) syncReplicationController(replicateController api.ReplicateController) {
	//fmt.Println("2224")
	podList := rm.client.ListPods(replicateController.Labels)
	diff := len(podList.Items) - replicateController.DesiredState.Replicas
	fmt.Println(diff)
	if diff < 0 {
		diff *= -1
		fmt.Printf("Too few replicas, creating %d\n", diff)
		for i := 0; i < diff; i++ {
			pod := api.Pod{
				JSONBase: api.JSONBase{
					ID: fmt.Sprintf("%x", rand.Int()),
				},
				DesiredState: replicateController.DesiredState.PodTemplate.DesiredState,
				Labels:       replicateController.DesiredState.PodTemplate.Labels,
			}
			rm.registry.CreatePod(pod, rm.scheduler.Schedule(pod))
		}
	} else if diff > 0 {
		fmt.Print("Too many replicas, deleting")
		for i := 0; i < diff; i++ {
			//rm.registry.DeleteTask()
		}
	}
}
