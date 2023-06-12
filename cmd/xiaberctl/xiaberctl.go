package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/xiaberctl"
	"io/ioutil"
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
			response, _ := client.Do(req)
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			println(string(body))
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
