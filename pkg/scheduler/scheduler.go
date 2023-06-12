package scheduler

import (
	"github.com/learnk8s/xiabernetes/pkg/types"
	"math/rand"
)

type Scheduler interface {
	Schedule(pod types.Pod) string
}

type RandomScheduler struct {
	nodes []string
}

func MakeRandomScheduler(nodes []string) *RandomScheduler {
	return &RandomScheduler{nodes: nodes}
}

func (rs *RandomScheduler) Schedule(pod types.Pod) string {
	return rs.nodes[rand.Intn(len(rs.nodes))]
}
