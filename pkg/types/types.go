package types

type JSONBase struct {
	ID string `json:"id,omitempty"`
}
type Task struct {
	JSONBase
	Labels       map[string]string `json:"labels,omitempty"`
	DesiredState TaskState         `json:"desiredState,omitempty"`
}

type TaskState struct {
	Manifest ContainerManifest `json:"manifest,omitempty"`
}

type ContainerManifest struct {
	Containers []Container `json:"containers,omitempty"`
}
type Container struct {
	Image string `json:"image,omitempty"`
	Ports []Port `json:"ports,omitempty"`
}

type Port struct {
	ContainerPort int `json:"containerPort,omitempty"`
	HostPort      int `json:"hostPort,omitempty"`
}
