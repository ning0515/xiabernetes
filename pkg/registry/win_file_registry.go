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

func (w *WinRegistry) CreateTask(task types.Task) {
	data, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll("../../storagepath/task", 0755)
	os.WriteFile("../../storagepath/task/"+task.ID+".txt", data, 0660)
}
