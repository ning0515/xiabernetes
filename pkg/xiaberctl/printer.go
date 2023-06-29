package xiaberctl

import (
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"io"
	"strings"
)

type Printer interface {
	Print([]byte, io.Writer) error
}
type HumanReadablePrinter struct {
}

var podColumn = []string{"Name", "Image", "Label"}
var controllerColumn = []string{"Name", "replicas", "Label"}
var statusColumn = []string{"Status"}

func PrintHeader(columnNames []string, w io.Writer) {
	fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))
}

func (hp *HumanReadablePrinter) PrintPodList(podList *api.PodList, w io.Writer) {
	for _, v := range podList.Items {
		fmt.Fprintf(w, "%s\t%v\t%s\n", v.ID, v.DesiredState.Manifest, v.Labels)
	}
}
func (hp *HumanReadablePrinter) PrintControllerList(controllers *api.ReplicateControllerList, w io.Writer) {
	for _, v := range controllers.Items {
		fmt.Fprintf(w, "%s\t%v\t%s\n", v.ID, v.DesiredState.Replicas, v.Labels)
	}
}

func (hp *HumanReadablePrinter) printStatus(status *api.Status, w io.Writer) {
	PrintHeader(statusColumn, w)
	fmt.Fprintf(w, "%v\n", status.Status)
}
func (hp *HumanReadablePrinter) Print(data []byte, w io.Writer) error {
	obj, err := api.Decode(data)
	if err != nil {
		return err
	}
	switch o := obj.(type) {
	case *api.PodList:
		PrintHeader(podColumn, w)
		hp.PrintPodList(o, w)
	case *api.ReplicateControllerList:
		PrintHeader(controllerColumn, w)
		hp.PrintControllerList(o, w)
	case *api.Status:
		hp.printStatus(o, w)
	default:
		fmt.Println("Wrong type")
	}
	return nil
}
