package registry

type TaskRegistry struct {
	storage TaskStorage
}

func MakeTaskRegistry(storage TaskStorage) *TaskRegistry {
	return &TaskRegistry{
		storage: storage,
	}
}

func (t *TaskRegistry) Create(name string) {
	t.storage.CreateTask(name)
}
