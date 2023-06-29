package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/labels"
)

type ControllerRegistryStorage struct {
	storage ControllerRegistry
}

func MakeControllerRegistryStorage(storage ControllerRegistry) *ControllerRegistryStorage {
	return &ControllerRegistryStorage{
		storage: storage,
	}
}

func (c *ControllerRegistryStorage) Create(controller interface{}) <-chan interface{} {
	newController := controller.(api.ReplicateController)
	c.storage.CreateController(newController)
	return apiserver.MakeAsync(func() interface{} {
		c.storage.CreateController(newController)
		return newController
	})
}

func (c *ControllerRegistryStorage) List(query labels.Query) interface{} {
	result := api.ReplicateControllerList{
		JSONBase: api.JSONBase{Kind: "cluster#replicationControllerList"},
		Items:    c.storage.ListController(query)}
	return result
}

func (c *ControllerRegistryStorage) Extract(data []byte) interface{} {
	controller := api.ReplicateController{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &controller)
	fmt.Printf("in Extract:\n %v\n", controller)
	return controller
}
