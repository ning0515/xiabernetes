package master

import (
	"github.com/learnk8s/xiabernetes/pkg/labels"
	"github.com/learnk8s/xiabernetes/pkg/registry"
	"sync"
	"time"
)

type PodCache struct {
	pods    registry.PodRegistry
	podInfo map[string]interface{}
	period  time.Duration
	podLock sync.Mutex
}

func NewPodCache(pods registry.PodRegistry, period time.Duration) *PodCache {
	return &PodCache{
		pods:    pods,
		period:  period,
		podInfo: map[string]interface{}{},
	}
}

func (p *PodCache) UpdateAllContainers() {
	pods := p.pods.ListPod(labels.Everything())
	for _, pod := range pods {
		p.podInfo[pod.ID] = pod
	}
}
