package main

import (
	"flag"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/xiaberctl"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	address  = flag.String("a", "http://127.0.0.1:8001", "Apiserver's endpoint")
	file     = flag.String("f", "", "The path of the config file")
	selector = flag.String("l", "", "label")
)

func usage() {
	log.Fatal("Usage:xiaberctl -a <address> [-f file.json][-p <hostPort>:<containerPort> <method> <path>]")
}

func readConfig(storage string) []byte {
	if len(*file) == 0 {
		log.Fatal("Need config file (-c)")
	}
	data, err := os.ReadFile(*file)
	if err != nil {
		log.Fatalf("Unable to read %v: %#v\n", *file, err)
	}
	data = xiaberctl.ToWireFormat(data, storage)
	if err != nil {
		log.Fatalf("Error parsing %v as an object for %v: %#v\n", *file, storage, err)
	}
	log.Printf("Parsed config file successfully; sending:\n%v\n", string(data))
	return data
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		usage()
	}
	printer := &xiaberctl.HumanReadablePrinter{}
	method := flag.Arg(0)
	storage := strings.Trim(flag.Arg(1), "/")
	url := *address + "/" + storage
	switch method {
	case "list":
		{
			if len(*selector) > 0 {
				url = url + "?labels=" + *selector
			}
			req, _ := http.NewRequest("GET", url, nil)
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()
			body, _ := io.ReadAll(response.Body)
			fmt.Println(string(body))
			printer.Print(string(body), os.Stdout)
		}
	case "create":
		{
			req := xiaberctl.RequestWithBody(readConfig(storage), url, "POST")
			client := &http.Client{}
			client.Do(req)
		}
	}
	//println(*address)
	//println(flag.NArg())
}
