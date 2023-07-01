package main

import (
	"flag"
	"fmt"
	"github.com/learnk8s/golang/glog"
	"github.com/learnk8s/xiabernetes/pkg/client"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"github.com/learnk8s/xiabernetes/pkg/xiaberctl"
	"os"
	"strings"
)

var (
	address  = flag.String("a", "http://127.0.0.1:8001", "Apiserver's endpoint")
	file     = flag.String("f", "", "The path of the config file")
	selector = flag.String("l", "", "label")
)

func usage() {
	glog.Info("Usage:xiaberctl -a <address> [-f file.json][-p <hostPort>:<containerPort> <method> <path>]")
}

func readConfig(storage string) []byte {
	if len(*file) == 0 {
		glog.Fatal("Need config file (-c)")
	}
	data, err := os.ReadFile(*file)
	if err != nil {
		glog.Fatalf("Unable to read %v: %#v\n", *file, err)
	}
	data = xiaberctl.ToWireFormat(data, storage)
	if err != nil {
		glog.Fatalf("Error parsing %v as an object for %v: %#v\n", *file, storage, err)
	}
	glog.Infof("Parsed config file successfully; sending:\n%v\n", string(data))
	return data
}

func main() {
	flag.Parse()
	util.InitLogs()
	defer util.FlushLogs()
	if flag.NArg() < 2 {
		usage()
		os.Exit(1)
	}
	printer := &xiaberctl.HumanReadablePrinter{}
	method := flag.Arg(0)
	storage := strings.Trim(flag.Arg(1), "/")
	verb := ""
	switch method {
	case "get", "list":
		verb = "GET"
	case "create":
		verb = "POST"
	}
	s := client.New(*address)
	r := s.Verb(verb).
		Path(storage).
		Query(*selector)
	if method == "create" {
		r.Body(readConfig(storage))
	}
	response, _ := r.Do()
	fmt.Println(string(response))
	printer.Print(response, os.Stdout)

	//switch method {
	//case "list":
	//	{
	//		if len(*selector) > 0 {
	//			url = url + "?labels=" + *selector
	//		}
	//		req, _ := http.NewRequest("GET", url, nil)
	//		client := &http.Client{}
	//		response, err := client.Do(req)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		defer response.Body.Close()
	//		body, _ := io.ReadAll(response.Body)
	//		fmt.Println(string(body))
	//		printer.Print(string(body), os.Stdout)
	//	}
	//case "create":
	//	{
	//		req := xiaberctl.RequestWithBody(readConfig(storage), url, "POST")
	//		client := &http.Client{}
	//		client.Do(req)
	//	}
	//}
	//println(*address)
	//println(flag.NArg())
}
