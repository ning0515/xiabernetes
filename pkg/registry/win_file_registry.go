package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/labels"
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

func (w *WinRegistry) CreatePod(pod types.Pod, node string) {
	data, err := json.MarshalIndent(pod, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	dir := "../../storagepath/hosts/" + node + "/pod/"
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+pod.ID+".txt", data, 0660)
}

func (w *WinRegistry) ListPod(label labels.Query) []types.Pod {
	pods := []types.Pod{}
	dir := "../../storagepath/hosts/"
	podList := ListFile(dir)
	for _, v := range podList {
		pod := types.Pod{}
		json.Unmarshal(v, &pod)
		if label.Matches(labels.Set(pod.Labels)) {
			pods = append(pods, pod)
		}
	}

	fmt.Printf("%v\n", pods)
	return pods
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

func (w *WinRegistry) ListController(label labels.Query) []types.ReplicateController {
	controllers := []types.ReplicateController{}
	dir := "../../storagepath/controllers/"
	controllerList := ListFile(dir)
	for _, v := range controllerList {
		controller := types.ReplicateController{}
		json.Unmarshal(v, &controller)
		if label.Matches(labels.Set(controller.Labels)) {
			controllers = append(controllers, controller)
		}
	}
	fmt.Printf("%v\n", controllers)
	return controllers
}

func ListFile(dir string) map[string][]byte {
	txtList := make(map[string][]byte, 10)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("%v", path)
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

func LabelsMatch(pod types.Pod, label *map[string]string) bool {
	if label == nil {
		return true
	}
	for key, value := range *label {
		if !LabelMatch(pod, key, value) {
			return false
		}
	}
	return true
}

func LabelMatch(pod types.Pod, queryKey, queryValue string) bool {
	for key, value := range pod.Labels {
		if queryKey == key && queryValue == value {
			return true
		}
	}
	return false
}
