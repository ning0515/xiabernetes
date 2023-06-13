package xiaberctl

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"io"
	"strings"
)

type Printer interface {
	Print(string, io.Writer) error
}
type HumanReadablePrinter struct {
}

var podColumn = []string{"Name", "Image", "Label"}
var controllerColumn = []string{"Name", "replicas", "Label"}

func PrintHeader(columnNames []string, w io.Writer) {
	fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))
}

func (hp *HumanReadablePrinter) PrintPod(data string, w io.Writer) {
	var pods types.PodList
	json.Unmarshal([]byte(data), &pods)
	for _, v := range pods.Items {
		fmt.Fprintf(w, "%s\t%v\t%s\n", v.ID, v.DesiredState.Manifest, v.Labels)
	}
}
func (hp *HumanReadablePrinter) PrintController(data string, w io.Writer) {
	var controllers types.ReplicateControllerList
	json.Unmarshal([]byte(data), &controllers)
	for _, v := range controllers.Items {
		fmt.Fprintf(w, "%s\t%v\t%s\n", v.ID, v.DesiredState.Replicas, v.Labels)
	}
}
func (hp *HumanReadablePrinter) Print(data string, w io.Writer) error {
	var obj interface{}
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return err
	}
	if _, contains := obj.(map[string]interface{})["kind"]; !contains {
		return fmt.Errorf("Unexpected object with no 'kind' field: %s", data)
	}
	kind := obj.(map[string]interface{})["kind"].(string)
	switch kind {
	case "cluster#podList":
		PrintHeader(podColumn, w)
		hp.PrintPod(data, w)
	case "cluster#replicationControllerList":
		PrintHeader(controllerColumn, w)
		hp.PrintController(data, w)
	}
	return nil
}
