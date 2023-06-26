package cache

import "sync"

type FIFO struct {
	lock  sync.RWMutex
	cond  sync.Cond
	items map[string]interface{}
	queue []string
}

func (f *FIFO) Add(ID string, obj interface{}) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.items[ID] = obj
	f.queue = append(f.queue, ID)
	f.cond.Broadcast()
}
func (f *FIFO) Update(ID string, obj interface{}) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.items[ID] = obj
	f.queue = append(f.queue, ID)
	f.cond.Broadcast()
}
func (f *FIFO) Delete(ID string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	delete(f.items, ID)
}
func (f *FIFO) Get(ID string) (obj interface{}) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	obj = f.items[ID]
	f.cond.Broadcast()
	return
}
func (f *FIFO) List() []interface{} {
	f.lock.RLock()
	defer f.lock.RUnlock()
	list := make([]interface{}, 0, len(f.items))
	for _, item := range f.items {
		list = append(list, item)
	}
	return list
}

func (f *FIFO) Pop() interface{} {
	f.lock.Lock()
	defer f.lock.Unlock()
	for {
		for len(f.queue) == 0 {
			f.cond.Wait()
		}
		id := f.queue[0]
		f.queue = f.queue[1:]
		item := f.items[id]
		delete(f.items, id)
		return item
	}
}

func NewFIFO() *FIFO {
	f := &FIFO{
		items: map[string]interface{}{},
		queue: []string{},
	}
	f.cond.L = &f.lock
	return f
}
