package registry

import "github.com/learnk8s/xiabernetes/pkg/types"

type TaskStorage interface {
	CreateTask(task types.Task, node string)
	ListTask(*map[string]string) []types.Task
}

type ControllerStorage interface {
	CreateController(replicateController types.ReplicateController)
	ListController() []types.ReplicateController
}
