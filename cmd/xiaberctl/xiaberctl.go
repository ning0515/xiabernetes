package main

import (
	"flag"
	"github.com/learnk8s/xiabernetes/pkg/xiaberctl"
	"log"
	"net/http"
	"os"
)

var address *string = flag.String("a", "127.0.0.1:8000", "Apiserver's endpoint")
var file *string = flag.String("f", "", "The path of the config file")

func usage() {
	log.Fatal("Usage:xiaberctl -a <address> [-f file.json][-p <hostPort>:<containerPort> <method> <path>]")
}
func main() {
	flag.Parse()
	//println(*file)
	data, err := os.ReadFile(*file)
	//println(string(data))
	if err != nil {
		println(err)
		return
	}

	req := xiaberctl.RequestWithBody(data, "POST")
	client := &http.Client{}
	client.Do(req)
	//config, _ := json.Marshal(data)
	//type aaa interface{}
	//var bbb aaa
	//fmt.Printf("%v\n", string(config))
	//json.Unmarshal(config, &bbb)
	//fmt.Println(bbb)

}
