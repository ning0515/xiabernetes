package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiServer struct {
	storage map[string]RESTStorage
}

type RESTStorage interface {
	Create(interface{})
	Extract(data []byte) interface{}
}

func New(storage map[string]RESTStorage) *ApiServer {
	return &ApiServer{
		storage: storage,
	}
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	path := r.URL.Path
	fmt.Println(path)
	resource := strings.Split(path, "/")
	fmt.Println(resource[1])
	//w.Write([]byte(resource[1]))
	data, _ := ioutil.ReadAll(r.Body)
	task := s.storage[resource[1]].Extract(data)
	req, _ := json.MarshalIndent(s.storage[resource[1]].Extract(data), "", "  ")
	fmt.Printf("%v", string(req))
	s.storage[resource[1]].Create(task)
}
