package registry

type TaskStorage interface {
	CreateTask(name string)
}
