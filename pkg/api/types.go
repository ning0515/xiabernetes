package api

type JSONBase struct {
	Kind string `json:"kind,omitempty"`
	ID   string `json:"id,omitempty"`
}

type PodList struct {
	JSONBase
	Items []Pod `json:"items" yaml:"items,omitempty"`
}

type Pod struct {
	JSONBase
	Labels       map[string]string `json:"labels,omitempty"`
	DesiredState PodState          `json:"desiredState,omitempty"`
}

type PodState struct {
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

type ReplicateController struct {
	JSONBase
	Labels       map[string]string        `json:"labels,omitempty"`
	DesiredState ReplicateControllerState `json:"desiredState,omitempty"`
}

type ReplicateControllerList struct {
	JSONBase
	Items []ReplicateController `json:"items,omitempty" yaml:"items,omitempty"`
}

type ReplicateControllerState struct {
	Replicas      int               `json:"replicas,omitempty"`
	ReplicasInSet map[string]string `json:"replicasInSet,omitempty"`
	PodTemplate   PodTemplate       `json:"podTemplate,omitempty"`
}

type PodTemplate struct {
	DesiredState PodState          `json:"desiredState,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
}

type Status struct {
	JSONBase `json:",inline" yaml:",inline"`
	// One of: "success", "failure", "working" (for operations not yet completed)
	// TODO: if "working", include an operation identifier so final status can be
	// checked.
	Status  string `json:"status,omitempty" yaml:"status,omitempty"`
	Details string `json:"details,omitempty" yaml:"details,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusFailure = "failure"
	StatusWorking = "working"
)

type ServerOp struct {
	JSONBase `yaml:",inline" json:",inline"`
}

// Operation list, as delivered to API clients.
type ServerOpList struct {
	JSONBase `yaml:",inline" json:",inline"`
	Items    []ServerOp `yaml:"items,omitempty" json:"items,omitempty"`
}
