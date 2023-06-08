package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiServer struct {
	storage map[string]RESTStorage
}

type RESTStorage interface {
	Create(name string)
	Extract(data []byte) interface{}
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path
	//resource := strings.Split(path, "/")
	//w.Write([]byte(resource[1]))
	//s.storage["tasks"].Extract(r.GetBody)
	task, _ := ioutil.ReadAll(r.Body)
	req, _ := json.MarshalIndent(s.storage["tasks"].Extract(task), "", "  ")
	fmt.Printf("%v", string(req))
}

func New(storage map[string]RESTStorage) *ApiServer {
	return &ApiServer{
		storage: storage,
	}
}
