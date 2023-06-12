package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func (w *WinRegistry) ListTask() []types.Task {
	tasks := []types.Task{}
	dir := "../../storagepath/hosts/"
	taskList := ListFile(dir)
	for _, v := range taskList {
		task := types.Task{}
		json.Unmarshal(v, &task)
		tasks = append(tasks, task)
	}

	fmt.Printf("%v\n", tasks)
	return tasks
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

func (w *WinRegistry) ListController() []types.ReplicateController {
	controllers := []types.ReplicateController{}
	dir := "../../storagepath/controllers/"
	controllerList := ListFile(dir)
	for _, v := range controllerList {
		controller := types.ReplicateController{}
		json.Unmarshal(v, &controller)
		controllers = append(controllers, controller)
	}

	fmt.Printf("%v\n", controllers)
	return controllers
}

func ListFile(dir string) map[string][]byte {
	txtList := make(map[string][]byte, 10)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(strings.LastIndexAny(path, "\\"))
			index := strings.LastIndexAny(path, "\\")
			ID := path[index+1 : len(path)-4]
			txtList[ID] = data
			//fmt.Printf("%s", txtList)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("错误：%v\n", err)
	}
	return txtList
}
