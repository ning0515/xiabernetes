package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/labels"
)

type ControllerRegistry struct {
	storage ControllerStorage
}

func MakeControllerRegistry(storage ControllerStorage) *ControllerRegistry {
	return &ControllerRegistry{
		storage: storage,
	}
}

func (c *ControllerRegistry) Create(controller interface{}) <-chan interface{} {
	newController := controller.(api.ReplicateController)
	c.storage.CreateController(newController)
	return apiserver.MakeAsync(func() interface{} {
		c.storage.CreateController(newController)
		return newController
	})
}

func (c *ControllerRegistry) List(query labels.Query) interface{} {
	result := api.ReplicateControllerList{
		JSONBase: api.JSONBase{Kind: "cluster#replicationControllerList"},
		Items:    c.storage.ListController(query)}
	return result
}

func (c *ControllerRegistry) Extract(data []byte) interface{} {
	controller := api.ReplicateController{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &controller)
	fmt.Printf("in Extract:\n %v\n", controller)
	return controller
}
