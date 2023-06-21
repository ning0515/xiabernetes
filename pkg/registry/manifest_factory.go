package registry

import (
	. "github.com/learnk8s/xiabernetes/pkg/api"
	"strings"
)

type ManifestFactory interface {
	// Make a container object for a given task, given the machine that the task is running on.
	MakeManifest(pod Pod) ContainerManifest
}

type BasicManifestFactory struct {
}

func (b *BasicManifestFactory) MakeManifest(pod Pod) ContainerManifest {
	for i, _ := range pod.DesiredState.Manifest.Containers {
		pod.DesiredState.Manifest.ID = pod.ID
		pod.DesiredState.Manifest.Containers[i].Name = strings.Trim(pod.DesiredState.Manifest.Containers[i].Image, "dockerfile/")
	}
	return pod.DesiredState.Manifest
}
