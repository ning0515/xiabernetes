package registry

import (
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/labels"
)

type PodStorage interface {
	CreatePod(api.Pod, string)
	ListPod(query labels.Query) []api.Pod
}

type ControllerStorage interface {
	CreateController(api.ReplicateController)
	ListController(query labels.Query) []api.ReplicateController
}
