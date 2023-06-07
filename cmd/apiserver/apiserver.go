package main

import (
	"github.com/learnk8s/xiabernetes/pkg/apiserver"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"net/http"
)

func main() {
	winRegistry := registry.MakeWinRegistry()
	storage := map[string]apiserver.RESTStorage{
		"tasks": registry.MakeTaskRegistry(winRegistry),
	}
	s := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: apiserver.New(storage),
	}
	s.ListenAndServe()
}
