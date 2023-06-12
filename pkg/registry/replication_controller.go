package registry

import "github.com/learnk8s/xiabernetes/pkg/apiserver"

type ReplicationManager struct {
	apiServer apiserver.ApiServer
	registry  WinRegistry
}
