package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type WinRegistry struct {
	manifestFactory ManifestFactory
}

func MakeWinRegistry() *WinRegistry {
	return &WinRegistry{manifestFactory: &BasicManifestFactory{}}
}

func (w *WinRegistry) CreatePod(pod api.Pod, node string) {
	w.runPod(pod, node)
}

func (w *WinRegistry) runPod(pod api.Pod, machine string) {
	manifests := w.LoadManifests(machine)
	data, err := json.MarshalIndent(pod, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	dir := "../../storagepath/hosts/" + machine + "/pod/"
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+pod.ID+".txt", data, 0660)
	manifest := w.manifestFactory.MakeManifest(pod)
	fmt.Printf("runPod manifest=%#v", manifest)
	manifests = append(manifests, manifest)
	w.updateManifests(machine, manifests)
}

func (w *WinRegistry) LoadManifests(machine string) []api.ContainerManifest {
	var manifests []api.ContainerManifest
	dir := "../../storagepath/hosts/" + machine + "/xiaberlet/"
	data, err := os.ReadFile(dir + "value.txt")
	if err != nil {
		fmt.Printf("err load manifest,err=%v", err)
		manifests = []api.ContainerManifest{}
	}
	json.Unmarshal(data, &manifests)
	return manifests
}
func (w *WinRegistry) updateManifests(machine string, manifests []api.ContainerManifest) {
	dir := "../../storagepath/hosts/" + machine + "/xiaberlet/"
	data, err := json.MarshalIndent(manifests, "", "	")
	if err != nil {
		fmt.Printf("error update manifests,err = %v", err)
	}
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+"value.txt", data, 0660)
}
func (w *WinRegistry) ListPod(label labels.Query) []api.Pod {
	pods := []api.Pod{}
	dir := "../../storagepath/hosts/"
	podList := ListFile(dir)
	for _, v := range podList {
		pod := api.Pod{}
		json.Unmarshal(v, &pod)
		if label.Matches(labels.Set(pod.Labels)) {
			pods = append(pods, pod)
		}
	}
	fmt.Printf("%v\n", pods)
	return pods
}

func (w *WinRegistry) CreateController(controller api.ReplicateController) {
	data, err := json.MarshalIndent(controller, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	dir := "../../storagepath/controllers/"
	os.MkdirAll(dir, 0755)
	os.WriteFile(dir+controller.ID+".txt", data, 0660)
}

func (w *WinRegistry) ListController(label labels.Query) []api.ReplicateController {
	controllers := []api.ReplicateController{}
	dir := "../../storagepath/controllers/"
	controllerList := ListFile(dir)
	for _, v := range controllerList {
		controller := api.ReplicateController{}
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
			if !strings.Contains(path, "xiaberlet") {
				data, err := os.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}
				index := strings.LastIndexAny(path, "\\")
				ID := path[index+1 : len(path)-4]
				txtList[ID] = data
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("错误：%v\n", err)
	}
	return txtList
}

func LabelsMatch(pod api.Pod, label *map[string]string) bool {
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

func LabelMatch(pod api.Pod, queryKey, queryValue string) bool {
	for key, value := range pod.Labels {
		if queryKey == key && queryValue == value {
			return true
		}
	}
	return false
}
