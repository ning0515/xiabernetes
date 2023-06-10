package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ApiServer struct {
	storage map[string]RESTStorage
}

type RESTStorage interface {
	Create(interface{})
	Extract(data []byte) interface{}
	List()
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
	switch r.Method {
	case "GET":
		{
			s.storage[resource[1]].List()
		}
	case "POST":
		{
			data, _ := io.ReadAll(r.Body)
			object := s.storage[resource[1]].Extract(data)
			req, _ := json.MarshalIndent(s.storage[resource[1]].Extract(data), "", "  ")
			fmt.Printf("%v", string(req))
			s.storage[resource[1]].Create(object)
		}
	}

}
