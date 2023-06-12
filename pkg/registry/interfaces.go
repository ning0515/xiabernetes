package registry

import "github.com/learnk8s/xiabernetes/pkg/types"

type PodStorage interface {
	CreatePod(pod types.Pod, node string)
	ListPod(*map[string]string) []types.Pod
}

type ControllerStorage interface {
	CreateController(replicateController types.ReplicateController)
	ListController() []types.ReplicateController
}
