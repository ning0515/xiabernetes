package main

import (
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"github.com/learnk8s/xiabernetes/pkg/xiaberlet"
	"time"
)

func main() {
	myXiaberlet := xiaberlet.Xiaberlet{
		FileRegistry:  registry.WinRegistry{},
		SyncFrequency: 10 * time.Second,
	}
	myXiaberlet.RunXiaberlet()
}
