package apiserver

import (
	"net/http"
	"strings"
)

type ApiServer struct {
	storage map[string]RESTStorage
}

type RESTStorage interface {
	Create(name string)
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	resource := strings.Split(path, "/")
	w.Write([]byte(resource[1]))
	s.storage["tasks"].Create(resource[1])
}

func New(storage map[string]RESTStorage) *ApiServer {
	return &ApiServer{
		storage: storage,
	}
}
