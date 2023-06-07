package registry

import "os"

type WinRegistry struct {
}

func MakeWinRegistry() *WinRegistry {
	return &WinRegistry{}
}

func (w *WinRegistry) CreateTask(name string) {
	os.Create(name + ".txt")
}
