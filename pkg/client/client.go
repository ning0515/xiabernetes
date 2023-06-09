package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"io"
	"net/http"
	"strings"
)

type ClientInterface interface {
	ListPods(map[string]string) api.PodList
	ListController() api.ReplicateControllerList
}

type StatusErr struct {
	Status api.Status
}

func (s *StatusErr) Error() string {
	return fmt.Sprintf("Status: %v (%#v)", s.Status.Status, s)
}

type Client struct {
	host       string
	httpClient *http.Client
}

func New(host string) *Client {
	return &Client{
		host: host,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func (c *Client) doRequest(request *http.Request) ([]byte, error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return body, err
	}
	// If the server gave us a status back, look at what it was.
	var status api.Status
	if err := api.DecodeInto(body, &status); err == nil && status.Status != "" {
		if status.Status == api.StatusSuccess {
			return body, nil
		}
		// "Working" requests need to be handled specially.
		// "Failed" requests are clearly just an error and it makes sense to return them as such.
		return nil, &StatusErr{status}
	}
	return body, err
}

func (c *Client) ListPods(label map[string]string) api.PodList {

	pods := api.PodList{}
	body, _ := c.Get().Path("pods").Query(LabelToString(label)).Do()
	err := json.Unmarshal(body, &pods)
	if err != nil {
		fmt.Println("json unmarshal error=", err)
	}
	println(string(body))
	return pods
	//url := c.Host + "/pods"
	//url = url + "?labels=" + LabelToString(label)
	//req, _ := http.NewRequest("GET", url, nil)
	//response, _ := c.HttpClient.Do(req)
	//defer response.Body.Close()
	//body, _ := io.ReadAll(response.Body)

}
func (c *Client) ListController() api.ReplicateControllerList {
	controllers := api.ReplicateControllerList{}
	body, _ := c.Get().Path("replicateController").Do()
	err := json.Unmarshal(body, &controllers)
	if err != nil {
		fmt.Println("json unmarshal error=", err)
	}
	println(string(body))
	return controllers
	//url := c.Host + "/replicateController"
	//req, _ := http.NewRequest("GET", url, nil)
	//response, _ := c.HttpClient.Do(req)
	//defer response.Body.Close()
	//body, _ := io.ReadAll(response.Body)
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
