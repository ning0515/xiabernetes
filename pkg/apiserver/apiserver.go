package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type ApiServer struct {
	storage map[string]RESTStorage
	ops     *Operations
}

type RESTStorage interface {
	Create(interface{}) <-chan interface{}
	Extract([]byte) interface{}
	List(query labels.Query) interface{}
}

func MakeAsync(fn func() interface{}) <-chan interface{} {
	channel := make(chan interface{}, 1)
	go func() {
		defer util.HandleCrash()
		channel <- fn()
	}()
	return channel
}
func New(storage map[string]RESTStorage) *ApiServer {
	return &ApiServer{
		storage: storage,
		ops:     NewOperations(),
	}
}

func parseTimeout(str string) time.Duration {
	if str != "" {
		timeout, err := time.ParseDuration(str)
		if err == nil {
			return timeout
		}
		fmt.Errorf("Failed to parse: %#v '%s'", err, str)
	}
	return 30 * time.Second
}

func (server *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sync := r.URL.Query().Get("sync") == "true"
	timeout := parseTimeout(r.URL.Query().Get("timeout"))
	fmt.Println(r.Method)
	path := r.URL.Path
	fmt.Println(path)

	resource := strings.Split(path, "/")
	fmt.Println(resource)
	//fmt.Println(len(resource[1:]))
	if resource[1] == "operations" {
		server.handleOperationRequest(resource[2:], w, r)
		return
	}
	switch r.Method {
	case "GET":
		{
			res := server.storage[resource[1]].List(labels.ParseQuery(r.URL.Query().Get("labels")))
			server.write(200, res, w)
		}
	case "POST":
		{
			data, _ := io.ReadAll(r.Body)
			object := server.storage[resource[1]].Extract(data)
			req, _ := json.MarshalIndent(server.storage[resource[1]].Extract(data), "", "  ")
			fmt.Printf("%v", string(req))
			fmt.Println("检验异步效果：提交了创建")
			out := server.storage[resource[1]].Create(object)
			fmt.Println("检验异步效果：主线程继续向下走")
			server.finishReq(out, sync, timeout, w)
		}
	}

}

func (server *ApiServer) write(statusCode int, object interface{}, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	output, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		log.Fatal(err)
		return
	}
	//fmt.Printf("222222333%v\n", output)
	w.Write(output)
}

func (server *ApiServer) finishReq(out <-chan interface{}, sync bool, timeout time.Duration, w http.ResponseWriter) {
	op := server.ops.NewOperation(out)
	if sync {
		op.WaitFor(timeout)
	}
	obj, finished := op.StatusOrResult()
	if finished {
		server.write(http.StatusOK, obj, w)
	} else {
		server.write(http.StatusAccepted, obj, w)
	}
}

func (server *ApiServer) handleOperationRequest(parts []string, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if len(parts) == 0 {
		list := server.ops.List()
		//fmt.Printf("22223%v\n", list)
		server.write(http.StatusOK, list, w)
		return
	}
	op := server.ops.Get(parts[0])
	//fmt.Print("%v", op)
	if op == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	obj, complete := op.StatusOrResult()
	if complete {
		server.write(http.StatusOK, obj, w)
	} else {
		server.write(http.StatusAccepted, obj, w)
	}
}
