package main

import (
	"flag"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/scheduler"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"net/http"
)

var (
	nodeList         util.StringList
	port             = flag.Uint("p", 8001, "Listing port")
	address          = flag.String("a", "127.0.0.1", "The address of api server")
	specifyScheduler = flag.String("scheduler", "random", "Specify a scheduler")
)

func init() {
	fmt.Sprint(1)
	flag.Var(&nodeList, "nodes", "List of nodes")
	fmt.Sprint(nodeList)
}
func main() {
	fmt.Sprint(nodeList)
	flag.Parse()
	var storage = map[string]apiserver.RESTStorage{}
	winRegistry := registry.MakeWinRegistry()
	//增加Scheduler就修改这里
	if *specifyScheduler == "random" {
		storage = map[string]apiserver.RESTStorage{
			"tasks": registry.MakeTaskRegistry(winRegistry, scheduler.MakeRandomScheduler(nodeList)),
		}
	} else {
		storage = map[string]apiserver.RESTStorage{
			"tasks": registry.MakeTaskRegistry(winRegistry, scheduler.MakeRandomScheduler(nodeList)),
		}
	}
	s := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: apiserver.New(storage),
	}
	s.ListenAndServe()
}
