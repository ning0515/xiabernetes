package registry

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/types"
)

type ControllerRegistry struct {
	storage ControllerStorage
}

func MakeControllerRegistry(storage ControllerStorage) *ControllerRegistry {
	return &ControllerRegistry{
		storage: storage,
	}
}

func (c *ControllerRegistry) Create(controller interface{}) {
	newController := controller.(types.ReplicateController)
	c.storage.CreateController(newController)
}

func (c *ControllerRegistry) List() {
	c.storage.ListController()
}

func (c *ControllerRegistry) Extract(data []byte) interface{} {
	controller := types.ReplicateController{}
	fmt.Printf("in data:\n %v\n", string(data))
	json.Unmarshal(data, &controller)
	fmt.Printf("in Extract:\n %v\n", controller)
	return controller
}
