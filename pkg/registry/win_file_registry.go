package registry

import (
	"encoding/json"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"log"
	"os"
)

type WinRegistry struct {
}

func MakeWinRegistry() *WinRegistry {
	return &WinRegistry{}
}

func (w *WinRegistry) CreateTask(task types.Task, node string) {
	data, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	dir := "../../storagepath/hosts/" + node + "/task/"
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+task.ID+".txt", data, 0660)
}

func (w *WinRegistry) CreateController(controller types.ReplicateController) {
	data, err := json.MarshalIndent(controller, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	dir := "../../storagepath/controllers/"
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+controller.ID+".txt", data, 0660)
}
