package xiaberctl

import (
	"github.com/learnk8s/xiabernetes/pkg/api"
	"reflect"
)

var storageToType = map[string]reflect.Type{
	"pods":                reflect.TypeOf(api.Pod{}),
	"replicateController": reflect.TypeOf(api.ReplicateController{}),
}

func ToWireFormat(data []byte, storage string) []byte {
	prototypeType, found := storageToType[storage]
	if !found {
		return nil
	}

	obj := reflect.New(prototypeType).Interface()
	err := api.DecodeInto(data, obj)
	if err != nil {
		return nil
	}
	data, err = api.Encode(obj)
	if err != nil {
		return nil
	}
	return data
}
