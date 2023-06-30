package api

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api/v1beta1"
	"gopkg.in/v1/yaml"
	"reflect"
)

type ConversionFunc func(input interface{}) (output interface{})

var versionMap = map[string]map[string]reflect.Type{}
var internalFuncs = map[string]ConversionFunc{}
var externalFuncs = map[string]ConversionFunc{}

func init() {
	AddKnownTypes("",
		PodList{},
		Pod{},
		ReplicateController{},
		ReplicateControllerList{},
		Status{},
		ServerOp{},
		ServerOpList{},
	)
	AddKnownTypes("v1beta1",
		v1beta1.PodList{},
		v1beta1.Pod{},
		v1beta1.ReplicateController{},
		v1beta1.ReplicateControllerList{},
		v1beta1.Status{},
		ServerOp{},
		ServerOpList{},
	)
}

func AddKnownTypes(version string, types ...interface{}) {
	knownTypes, found := versionMap[version]
	if !found {
		knownTypes = map[string]reflect.Type{}
		versionMap[version] = knownTypes
	}
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
	data, err = json.MarshalIndent(obj, "", "    ")
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
	knownTypes := versionMap[jsonBase.APIVersion]
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

func Decode(data []byte) (interface{}, error) {
	findKind := struct {
		Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
		APIVersion string `json:"apiVersion,omitempty"`
	}{}
	err := json.Unmarshal(data, &findKind)
	if err != nil {
		fmt.Errorf("Unmarshal error=%v", err)
		return nil, err
	}
	knownTypes := versionMap[findKind.APIVersion]
	objType, found := knownTypes[findKind.Kind]
	if !found {
		return nil, fmt.Errorf("%v is not a known type", findKind.Kind)
	}
	obj := reflect.New(objType).Interface()
	err = json.Unmarshal(data, obj)
	if err != nil {
		fmt.Errorf("Unmarshal error=%v", err)
		return nil, err
	}
	_, jsonBase, err := nameAndJSONBase(obj)
	if err != nil {
		return nil, err
	}
	// Don't leave these set. Track type with go's type.
	jsonBase.Kind = ""
	return obj, nil
}
func DecodeInto(data []byte, obj interface{}) error {
	fmt.Printf("before unmarshal %s\n", data)
	err := yaml.Unmarshal(data, obj)
	fmt.Printf("after unmarshal %#v\n", obj)
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
