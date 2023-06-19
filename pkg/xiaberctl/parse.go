package xiaberctl

import (
	"encoding/json"
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
	err := json.Unmarshal(data, obj)
	if err != nil {
		return nil
	}
	data, err = json.Marshal(obj)
	if err != nil {
		return nil
	}
	return data
}
