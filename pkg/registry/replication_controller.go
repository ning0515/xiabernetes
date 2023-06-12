package registry

import (
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"math/rand"
)

type ReplicationManager struct {
	registry  WinRegistry
	scheduler scheduler.Scheduler
}

func MakeReplicateManager(registry WinRegistry, scheduler scheduler.Scheduler) *ReplicationManager {
	return &ReplicationManager{
		registry:  registry,
		scheduler: scheduler,
	}
}
func (rm *ReplicationManager) Sync() {
	replicateControllers := rm.registry.ListController()
	for _, replicateController := range replicateControllers {
		rm.syncReplicationController(replicateController)
	}
}

func (rm *ReplicationManager) syncReplicationController(replicateController types.ReplicateController) {
	podList := rm.registry.ListPod(&replicateController.Labels)
	diff := len(podList) - replicateController.DesiredState.Replicas
	if diff < 0 {
		diff *= -1
		fmt.Printf("Too few replicas, creating %d\n", diff)
		for i := 0; i < diff; i++ {
			pod := types.Pod{
				JSONBase: types.JSONBase{
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
