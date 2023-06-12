package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/client"
)

var (
	apiServer = flag.String("server", "127.0.0.1:8001", "apiserver")
	registry  = flag.String("r", "win", "registry")
)

func main() {
	label := map[string]string{"name": "foo"}
	client := client.Client{
		Host: "http://" + *apiServer,
	}
	client.ListTasks(label)
}
