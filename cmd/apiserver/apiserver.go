package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/master"
	"github.com/learnk8s/xiabernetes/pkg/util"
)

var (
	nodeList         util.StringList
	port             = flag.Uint("p", 8001, "Listing port")
	address          = flag.String("address", "127.0.0.1", "The address of api server")
	apiPrefix        = flag.String("api_prefix", "/api/v1beta1", "The prefix for API requests on the server. Default '/api/v1beta1'")
	specifyScheduler = flag.String("scheduler", "random", "Specify a scheduler")
)

func init() {
	flag.Var(&nodeList, "nodes", "List of nodes")
}
func main() {
	flag.Parse()
	m := master.New(nodeList)
	m.Run(*address, *apiPrefix)
}
