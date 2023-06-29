package master

import (
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	"net/http"
)

type Master struct {
	podRegistry        registry.PodRegistry
	controllerRegistry registry.ControllerRegistry

	storage map[string]apiserver.RESTStorage
	minions []string

	ops *apiserver.Operations
}

func New(minions []string) *Master {
	m := &Master{
		podRegistry:        registry.MakeWinRegistry(),
		controllerRegistry: registry.MakeWinRegistry(),
		ops:                apiserver.NewOperations(),
	}
	m.init(minions)
	return m
}

func (m *Master) init(minions []string) {
	m.minions = minions
	m.storage = map[string]apiserver.RESTStorage{
		"pods":                registry.MakePodRegistryStorage(m.podRegistry, scheduler.MakeRandomScheduler(m.minions)),
		"replicateController": registry.MakeControllerRegistryStorage(m.controllerRegistry),
	}
}

func (m *Master) Run(address, apiPrefix string) {
	s := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: apiserver.New(m.storage),
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
