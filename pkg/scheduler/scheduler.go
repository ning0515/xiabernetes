package scheduler

import (
	"github.com/learnk8s/xiabernetes/pkg/api"
	"math/rand"
)

type Scheduler interface {
	Schedule(api.Pod) string
}

type RandomScheduler struct {
	nodes []string
}

func MakeRandomScheduler(nodes []string) *RandomScheduler {
	return &RandomScheduler{nodes: nodes}
}

func (rs *RandomScheduler) Schedule(pod api.Pod) string {
	return rs.nodes[rand.Intn(len(rs.nodes))]
}
