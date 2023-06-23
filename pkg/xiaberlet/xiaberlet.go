package xiaberlet

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Xiaberlet struct {
	FileRegistry  registry.WinRegistry
	SyncFrequency time.Duration
}

func (xl *Xiaberlet) RunXiaberlet() {
	winChannel := make(chan []api.ContainerManifest)
	go util.Forever(func() { xl.WatchWin(winChannel) }, 10*time.Second)
	xl.RunSyncLoop(winChannel)
}

func (xl *Xiaberlet) RunSyncLoop(winChannel <-chan []api.ContainerManifest) {
	var lastWin []api.ContainerManifest
	for {
		select {
		case manifests := <-winChannel:
			lastWin = manifests
		case <-time.After(xl.SyncFrequency):
		}
		manifests := append([]api.ContainerManifest{}, lastWin...)
		xl.SyncManifests(manifests)

	}
}

type SyncHandler interface {
	SyncManifests([]api.ContainerManifest) error
}

func (xl *Xiaberlet) SyncManifests(config []api.ContainerManifest) error {
	fmt.Printf("Desired:%#v\n", config)
	desired := map[string]bool{}
	for _, manifest := range config {
		for _, element := range manifest.Containers {
			var exists bool
			exists, foundName := xl.ContainerExists(manifest, element)
			fmt.Printf("syncManifests exists=%v\n", exists)
			if !exists {
				name := xl.RunContainer(&manifest, &element)
				desired[name] = true
			} else {
				fmt.Printf("found container %s\n", foundName)
			}
			desired[foundName] = true
		}
	}

	return nil
}

func (xl *Xiaberlet) RunContainer(manifest *api.ContainerManifest, container *api.Container) (name string) {
	name = fmt.Sprintf("%s--%s", container.Name, manifest.ID)
	dir := "../../storagepath/fakedockerpath/"
	data, err := json.Marshal(container)
	if err != nil {
		fmt.Printf("run container error ,err = %#v", err)
	}
	os.MkdirAll(dir, 0755)
	fmt.Printf("runContainer dir=%v\n", dir+name+".txt")
	err = os.WriteFile(dir+name+".txt", data, 0660)
	if err != nil {
		fmt.Printf("write error=%v\n", err)
	}
	return name
}
func (xl *Xiaberlet) ContainerExists(manifests api.ContainerManifest, container api.Container) (exists bool, foundName string) {
	containers := xl.ListContainers()
	for _, name := range containers {
		if name == container.Name+"--"+manifests.ID {
			return true, name
		}
	}
	return false, ""
}
func (xl *Xiaberlet) ListContainers() []string {
	result := []string{}
	dir := "../../storagepath/fakedockerpath/"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("%v", path)
			return err
		}
		if !info.IsDir() {
			index := strings.LastIndexAny(path, "\\")
			ID := path[index+1 : len(path)-4]
			result = append(result, ID)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("错误：%v\n", err)
	}
	return result
}

func (xl *Xiaberlet) WatchWin(changeChannel chan<- []api.ContainerManifest) {
	//	var manifests []api.ContainerManifest
	newData := xl.FileRegistry.LoadManifests("1.1.1.1")
	//manifests = append(manifests, newData...)
	changeChannel <- newData
}

type podWorkers struct {
	lock    sync.Mutex
	workers util.StringSet
}

func newPodWorkers() podWorkers {
	return podWorkers{
		workers: util.NewStringSet(),
	}
}

func (self *podWorkers) Run(podFullName string, action func()) {
	self.lock.Lock()
	defer self.lock.Unlock()

	// This worker is already running, let it finish.
	if self.workers.Has(podFullName) {
		return
	}
	self.workers.Insert(podFullName)

	// Run worker async.
	go func() {
		defer util.HandleCrash()
		action()

		self.lock.Lock()
		defer self.lock.Unlock()
		self.workers.Delete(podFullName)
	}()
}
