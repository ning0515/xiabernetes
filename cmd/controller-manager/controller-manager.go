package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/client"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"net/http"
	"time"
)

var (
	nodeList  util.StringList
	apiServer = flag.String("server", "127.0.0.1:8001", "apiserver")
	//registry  = flag.String("r", "win", "registry")
)

func init() {
	flag.Var(&nodeList, "nodes", "List of nodes")
}

func main() {
	flag.Parse()
	nodeList = append(nodeList, "1.1.1.1")
	label := map[string]string{"name": "foo"}
	client := client.Client{
		Host:       "http://" + *apiServer,
		HttpClient: &http.Client{},
	}
	client.ListPods(label)
	controllerManager := registry.MakeReplicateManager(*registry.MakeWinRegistry(), scheduler.MakeRandomScheduler(nodeList), client)

	go util.Forever(func() { controllerManager.Sync() }, 1*time.Second)
	select {}
}
