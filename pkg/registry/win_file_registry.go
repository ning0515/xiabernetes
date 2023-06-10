package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"log"
	"os"
	"path/filepath"
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

func (w *WinRegistry) ListTask() {
	dir := "../../storagepath/hosts/"
	ListFile(dir)
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

func (w *WinRegistry) ListController() {
	dir := "../../storagepath/controllers/"
	ListFile(dir)
}

func ListFile(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("错误：%v\n", err)
	}
}
