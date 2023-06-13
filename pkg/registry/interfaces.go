package registry

import "github.com/learnk8s/xiabernetes/pkg/types"

type PodStorage interface {
	CreatePod(types.Pod, string)
	ListPod(*map[string]string) []types.Pod
}

type ControllerStorage interface {
	CreateController(types.ReplicateController)
	ListController() []types.ReplicateController
}
