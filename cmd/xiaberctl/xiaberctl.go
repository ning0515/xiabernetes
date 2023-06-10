package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/xiaberctl"
	"log"
	"net/http"
	"os"
)

var (
	address = flag.String("a", "http://127.0.0.1:8000", "Apiserver's endpoint")
	file    = flag.String("f", "", "The path of the config file")
)

func usage() {
	log.Fatal("Usage:xiaberctl -a <address> [-f file.json][-p <hostPort>:<containerPort> <method> <path>]")
}
func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		usage()
	}
	method := flag.Arg(0)
	url := *address
	switch method {
	case "list":
		{
			url += flag.Arg(1)
			req, _ := http.NewRequest("GET", url, nil)
			client := &http.Client{}
			client.Do(req)
		}
	case "create":
		{
			url += flag.Arg(1)
			data, err := os.ReadFile(*file)
			if err != nil {
				log.Fatal(err)
				return
			}
			req := xiaberctl.RequestWithBody(data, url, "POST")
			client := &http.Client{}
			client.Do(req)
		}
	}
	println(*address)
	println(flag.NArg())
}
