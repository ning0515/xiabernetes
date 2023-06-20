package registry

import . "github.com/learnk8s/xiabernetes/pkg/api"

type ManifestFactory interface {
	// Make a container object for a given task, given the machine that the task is running on.
	MakeManifest(pod Pod) ContainerManifest
}

type BasicManifestFactory struct {
}

func (b *BasicManifestFactory) MakeManifest(pod Pod) ContainerManifest {
	for _, v := range pod.DesiredState.Manifest.Containers {
		v.ID = pod.ID
	}
	return pod.DesiredState.Manifest
}
