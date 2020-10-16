package store

import (
	"container/list"
	"sync"
)

//EmptyQueue
type emptyQueue struct {
	list *list.List
	mu   sync.Mutex
	emptyq
}
type emptyq interface {
	Push(*nodeStruct)
	Pop() (*nodeStruct, error)
}

//Data type this queue carries
type emptyLoc struct {
	treeI   uint32
	branchI uint32
	nPtr    *nodeStruct
}

func (e *emptyQueue) Push(n *nodeStruct) {
	e.mu.Lock()
	e.list.PushBack(n)
	e.mu.Unlock()
	return
}

func (e *emptyQueue) Pop() *nodeStruct {
	if e.list.Len() == 0 {
		return nil
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	n := e.list.Front()
	node := n.Value
	e.list.Remove(n)
	//I haven't found a better way than an assertion here, look into it again
	if node, ok := node.(*nodeStruct); ok {
		return node
	}
	//Nil == error
	return nil
}

//### COMPLETED QUEUE

//CompQueue stucture
type compQueue struct {
	list *list.List
	mu   sync.Mutex
	cq
}
type cq interface {
	New(interface{})
	Pop() interface{}
}

//New .. locks & adds a new task to the queue
func (cq *compQueue) New(task COMPLETED) {
	cq.mu.Lock()
	cq.list.PushBack(task)
	cq.mu.Unlock()
	return
}

//Pop .. locks & gets a new task
func (cq *compQueue) Pop() COMPLETED {
	cq.mu.Lock()
	c := cq.list.Front()
	comp := c.Value
	cq.list.Remove(c)
	cq.mu.Unlock()
	//I haven't found a better way than an assertion here, look into it again
	if comp, ok := comp.(COMPLETED); ok {
		return comp
	}
	//Nil == error
	return COMPLETED{}
}

//### TASK QUEUE

//TaskQueue stucture
type taskQueue struct {
	list *list.List
	mu   sync.Mutex
	tq
}
type tq interface {
	New(interface{})
	Pop() interface{}
}

//New .. locks & adds a new task to the queue
func (tq *taskQueue) New(task interface{}) {
	tq.mu.Lock()
	tq.list.PushBack(task)
	tq.mu.Unlock()
	return
}

//Pop .. locks & gets a new task
func (tq *taskQueue) Pop() interface{} {
	tq.mu.Lock()
	t := tq.list.Front()
	task := t.Value
	tq.list.Remove(t)
	tq.mu.Unlock()
	return task
}

//#### ERROR QUEUE

//ErrorQueue structure
type errorQueue struct {
	List *list.List
	sync.Mutex
	sync.Cond
	eq
}
type eq interface {
	New(interface{})
}

//New .. locks & adds a new error to the queue
func (eq *errorQueue) New(task interface{}) {
	eq.Lock()
	eq.List.PushBack(task)
	eq.Unlock()
	eq.Signal()
	return
}

//Pop .. locks & gets a new task
func (eq *errorQueue) Pop() interface{} {
	eq.Lock()
	e := eq.List.Front()
	err := e.Value
	eq.List.Remove(e)
	eq.Unlock()
	return err
}
