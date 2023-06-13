package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiServer struct {
	storage map[string]RESTStorage
}

type RESTStorage interface {
	Create(interface{})
	Extract([]byte) interface{}
	List(*url.URL) interface{}
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
			res := s.storage[resource[1]].List(r.URL)
			s.write(200, res, w)
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

func (server *ApiServer) write(statusCode int, object interface{}, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	output, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(output)
}
