package registry

import "github.com/learnk8s/xiabernetes/pkg/types"

type TaskStorage interface {
	CreateTask(task types.Task, node string)
}

type ControllerStorage interface {
	CreateController(replicateController types.ReplicateController)
}
