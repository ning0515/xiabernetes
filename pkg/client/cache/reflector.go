package cache

import (
	"github.com/learnk8s/xiabernetes/pkg/client"
	"reflect"
)

type Store interface {
	Add(ID string, obj interface{})
	Update(ID string, obj interface{})
	Delete(ID string)
	List() []interface{}
	Get(ID string) (item interface{}, exists bool)
}

type Reflector struct {
	client     *client.Client
	resource   string
	expectType reflect.Type
	store      Store
}

func NewReflector(client *client.Client, resource string, expectType interface{}, store Store) *Reflector {
	result := &Reflector{
		client:     client,
		resource:   resource,
		expectType: reflect.TypeOf(expectType),
		store:      store,
	}
	return result
}
