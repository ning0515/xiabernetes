package apiserver

import (
	"github.com/learnk8s/xiabernetes/pkg/types"
	"github.com/learnk8s/xiabernetes/pkg/util"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Operation struct {
	ID       string
	result   interface{}
	awaiting <-chan interface{}
	finished *time.Time
	lock     sync.Mutex
	notify   chan bool
}

// Operations tracks all the ongoing operations.
type Operations struct {
	// Access only using functions from atomic.
	lastID int64

	// 'lock' guards the ops map.
	lock sync.Mutex
	ops  map[string]*Operation
}

func NewOperations() *Operations {
	ops := &Operations{
		ops: map[string]*Operation{},
	}
	go util.Forever(func() { ops.expire(10 * time.Minute) }, 5*time.Minute)
	return ops
}

func (ops *Operations) NewOperation(from <-chan interface{}) *Operation {
	id := atomic.AddInt64(&ops.lastID, 1)
	op := &Operation{
		ID:       strconv.FormatInt(id, 10),
		awaiting: from,
		notify:   make(chan bool, 1),
	}
	go op.wait()
	go ops.insert(op)
	return op
}
func (op *Operation) wait() {
	defer util.HandleCrash()
	result := <-op.awaiting

	op.lock.Lock()
	defer op.lock.Unlock()
	op.result = result
	finished := time.Now()
	op.finished = &finished
	op.notify <- true
}

func (op *Operation) WaitFor(timeout time.Duration) {
	select {
	case <-time.After(timeout):
	case <-op.notify:
		// Re-send on this channel in case there are others
		// waiting for notification.
		op.notify <- true
	}
}
func (ops *Operations) List() types.ServerOpList {
	ops.lock.Lock()
	defer ops.lock.Unlock()

	ids := []string{}
	for id := range ops.ops {
		ids = append(ids, id)
	}
	sort.StringSlice(ids).Sort()
	ol := types.ServerOpList{}
	for _, id := range ids {
		ol.Items = append(ol.Items, types.ServerOp{JSONBase: types.JSONBase{ID: id}})
	}
	//fmt.Println(1111)
	//fmt.Print("%v", ol)
	return ol
}
func (ops *Operations) Get(id string) *Operation {
	ops.lock.Lock()
	defer ops.lock.Unlock()
	return ops.ops[id]
}
func (op *Operation) StatusOrResult() (description interface{}, finished bool) {
	op.lock.Lock()
	defer op.lock.Unlock()

	if op.finished == nil {
		return types.Status{
			Status:  types.StatusWorking,
			Details: op.ID,
		}, false
	}
	return op.result, true
}

func (ops *Operations) insert(op *Operation) {
	ops.lock.Lock()
	defer ops.lock.Unlock()
	ops.ops[op.ID] = op
}
func (ops *Operations) expire(maxAge time.Duration) {
	ops.lock.Lock()
	defer ops.lock.Unlock()
	keep := map[string]*Operation{}
	limitTime := time.Now().Add(-maxAge)
	for id, op := range ops.ops {
		if !op.expired(limitTime) {
			keep[id] = op
		}
	}
	ops.ops = keep

}

func (op *Operation) expired(lastTime time.Time) bool {
	op.lock.Lock()
	defer op.lock.Unlock()
	if op.finished == nil {
		return false
	}
	return op.finished.Before(lastTime)
}
