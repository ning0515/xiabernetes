package xiaberlet

import (
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"time"
)

type Xiaberlet struct {
	FileRegistry  registry.WinRegistry
	SyncFrequency time.Duration
}

func (xl *Xiaberlet) RunXiaberlet() {
	winChannel := make(chan []api.ContainerManifest)
	go util.Forever(func() { xl.WatchWin(winChannel) }, 10*time.Second)
	//xl.
}

func (xl *Xiaberlet) RunSyncLoop(winChannel <-chan []api.ContainerManifest) {
	//var lastWin []api.ContainerManifest
	//for {
	//	select {
	//	case manifests := <-winChannel:
	//		lastWin = manifests
	//	case <-time.After(xl.SyncFrequency):
	//	}
	//	manifests := append([]api.ContainerManifest{},lastWin...)
	//
	//}
}

type SyncHandler interface {
	SyncManifests([]api.ContainerManifest) error
}

func (xl *Xiaberlet) SyncManifests(config []api.ContainerManifest) error {
	fmt.Printf("Desired:")
	return nil
}

func (xl *Xiaberlet) WatchWin(changeChannel chan<- []api.ContainerManifest) {
	var manifests []api.ContainerManifest
	newData := xl.FileRegistry.ListPod(labels.ParseQuery(""))
	for _, v := range newData {
		manifests = append(manifests, v.DesiredState.Manifest)
	}
	changeChannel <- manifests
}
