package registry

import (
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/labels"
)

type PodRegistry interface {
	CreatePod(api.Pod, string)
	ListPod(query labels.Query) []api.Pod
}

type ControllerRegistry interface {
	CreateController(api.ReplicateController)
	ListController(query labels.Query) []api.ReplicateController
}
