package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/v1/yaml"
	"reflect"
)

var knownTypes = map[string]reflect.Type{}

func init() {
	AddKnownTypes(
		PodList{},
		Pod{},
		ReplicateController{},
		ReplicateControllerList{},
		Status{},
	)
}

func AddKnownTypes(types ...interface{}) {
	for _, obj := range types {
		t := reflect.TypeOf(obj)
		knownTypes[t.Name()] = t
	}
}
func Encode(obj interface{}) (data []byte, err error) {
	obj = checkPtr(obj)
	jsonBase, err := prepareEncode(obj)
	if err != nil {
		return nil, err
	}
	data, err = json.MarshalIndent(obj, "", "	")
	jsonBase.Kind = ""
	return data, err
}
func checkPtr(obj interface{}) interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		return obj
	}
	v2 := reflect.New(v.Type())
	v2.Elem().Set(v)
	return v2.Interface()
}

func prepareEncode(obj interface{}) (*JSONBase, error) {
	name, jsonBase, err := nameAndJSONBase(obj)
	if err != nil {
		return nil, err
	}
	if _, contains := knownTypes[name]; !contains {
		return nil, fmt.Errorf("struct %v won't be unmarshalable because it's not in knownTypes", name)
	}
	jsonBase.Kind = name
	return jsonBase, nil
}

func nameAndJSONBase(obj interface{}) (string, *JSONBase, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return "", nil, fmt.Errorf("expected pointer, but got %v", v.Type().Name())
	}
	v = v.Elem()
	name := v.Type().Name()
	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("expected struct, but got %v", name)
	}
	jsonBase := v.FieldByName("JSONBase")
	if !jsonBase.IsValid() {
		return "", nil, fmt.Errorf("struct %v lacks embedded JSON type", name)
	}
	return name, jsonBase.Addr().Interface().(*JSONBase), nil
}

func DecodeInto(data []byte, obj interface{}) error {
	err := yaml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	name, jsonBase, err := nameAndJSONBase(obj)
	if err != nil {
		return err
	}
	if jsonBase.Kind != "" && jsonBase.Kind != name {
		return fmt.Errorf("data had kind %v, but passed object was of type %v", jsonBase.Kind, name)
	}
	// Don't leave these set. Track type with go's type.
	jsonBase.Kind = ""
	return nil
}
