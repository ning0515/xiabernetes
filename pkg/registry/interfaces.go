package registry

import (
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/types"
)

type PodStorage interface {
	CreatePod(types.Pod, string)
	ListPod(query labels.Query) []types.Pod
}

type ControllerStorage interface {
	CreateController(types.ReplicateController)
	ListController(query labels.Query) []types.ReplicateController
}
