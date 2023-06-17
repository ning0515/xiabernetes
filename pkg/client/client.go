package client

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/types"
	"io"
	"net/http"
	"strings"
)

type ClientInterface interface {
	ListPods(map[string]string) types.PodList
	ListController() types.ReplicateControllerList
}

type Client struct {
	Host       string
	HttpClient *http.Client
}

func (c Client) ListPods(label map[string]string) types.PodList {
	pods := types.PodList{}
	url := c.Host + "/pods"
	url = url + "?labels=" + LabelToString(label)
	req, _ := http.NewRequest("GET", url, nil)
	response, _ := c.HttpClient.Do(req)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &pods)
	if err != nil {
		fmt.Println("json unmarshal error=", err)
	}
	println(string(body))
	return pods

}
func (c Client) ListController() types.ReplicateControllerList {
	controllers := types.ReplicateControllerList{}
	url := c.Host + "/replicateController"
	req, _ := http.NewRequest("GET", url, nil)
	response, _ := c.HttpClient.Do(req)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &controllers)
	if err != nil {
		fmt.Println("json unmarshal error=", err)
	}
	println(string(body))
	return controllers
}

func LabelToString(label map[string]string) string {
	labelSlice := make([]string, 0, len(label))
	for key, value := range label {
		labelSlice = append(labelSlice, key+"="+value)
	}
	return strings.Join(labelSlice, ",")
}

func StringToLabel(labelString string) map[string]string {
	label := map[string]string{}
	if len(labelString) == 0 {
		return label
	}
	parts := strings.Split(labelString, ",")
	for _, part := range parts {
		keyValue := strings.Split(part, "=")
		if len(keyValue) == 2 {
			label[keyValue[0]] = keyValue[1]
		} else {
			fmt.Printf("Wrong serch")
		}
	}
	return label
}
