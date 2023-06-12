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
	ListTasks(map[string]string) types.TaskList
}

type Client struct {
	Host       string
	httpClient *http.Client
}

func (c Client) ListTasks(label map[string]string) types.TaskList {
	tasks := types.TaskList{}
	url := c.Host + "/tasks"
	url = url + "?labels=" + LabelToString(label)
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, tasks)
	println(string(body))
	return types.TaskList{}

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
