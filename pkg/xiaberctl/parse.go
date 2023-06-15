package xiaberctl

import (
	"encoding/json"
	"reflect"
)
import "github.com/learnk8s/xiabernetes/pkg/types"

var storageToType = map[string]reflect.Type{
	"pods":                reflect.TypeOf(types.Pod{}),
	"replicateController": reflect.TypeOf(types.ReplicateController{}),
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
