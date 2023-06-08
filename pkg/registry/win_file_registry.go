package registry

import "os"

type WinRegistry struct {
}

func MakeWinRegistry() *WinRegistry {
	return &WinRegistry{}
}

func (w *WinRegistry) CreateTask(name string) {
	os.MkdirAll("../../storagepath/task", 0755)
	os.WriteFile("../../storagepath/task/"+name+".txt", []byte("Hello, Gophers!"), 0660)
}
