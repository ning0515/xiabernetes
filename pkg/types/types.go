package types

type JSONBase struct {
	ID string `id,omitempty`
}
type Task struct {
	labels       map[string]string `labels,omitempty`
	desiredState TaskState         `desiredState,omitempty`
}

type TaskState struct {
	manifest ContainerManifest `manifest,omitempty`
}

type ContainerManifest struct {
	containers []Container `containers,omitempty`
}
type Container struct {
	image string `image,omitempty`
	ports []Port `ports,omitempty`
}

type Port struct {
	containerPort int `containerPort`
	hostPort      int `hostPort`
}
