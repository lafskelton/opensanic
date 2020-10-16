package store

import (
	"container/list"
	"fmt"
	"time"
)

const (
	maxArrSize = 11111
)

//DocumentStore ..
type DocumentStore struct {
	Name string
	//Options
	Opt Options
	//D: Interface of methods
	D
	//nodes: Stack of nodes
	Tree *SanicBST
	//Worker Pool
	Pool *WorkerPool
}

//Options ..
type Options struct {
	MaxNodeCount uint64
}

//D methods
type D interface {
	//Public
	//
	//Start store
	Start()
	//TODO Status of a store: Avg read/write times, error rate, size, etc.
	//Status()
	//Get a node
	Get(uint32) (string, error)
	//Set a node
	Set(uint32, string) error
	//Update a node
	Update(uint32, string) error
	//Delete a node
	Delete(uint32) error
	//
}

//Start a document store
func (d *DocumentStore) Start(name string) error {
	//
	//	#### TODO
	// 	Check avaliable memory, each node has a maximum capacity of 12mb,
	//	ensure the tree fits when fully expanded with X overhead
	//
	d.Name = name
	//Init new tree
	d.Tree = new(SanicBST)
	d.Tree.Init()
	//Start task pool worker
	d.Pool = new(WorkerPool)
	d.Pool.Errs.List = list.New()
	d.Pool.Task.Queue.list = list.New()
	d.Pool.Finished.Queue.list = list.New()
	//Assign memory worker callback
	d.Pool.collapse = func() error {
		return nil
	}
	//Give the pool a pointer to it's owner
	d.Pool.store = d
	//pass locker to worker conditionals
	d.Pool.Task.NewTask.L = &d.Pool.Task.Mu
	d.Pool.Errs.L = &d.Pool.Errs.Mutex
	d.Pool.Finished.NewComplete.L = &d.Pool.Finished.Mu
	//Finish Everything
	time.Sleep(time.Millisecond * 100)
	//Start Workers
	go d.Pool.StartAll()
	return nil
}

// Get //
// Get a value into the document tree
func (d *DocumentStore) Get(key uint32) (string, error) {
	val, err := d.Tree.GET(key)
	if err != nil {
		return "", nil
	}
	//Convert byte to string, there must be a better way to do this...
	return string(val), nil
}

// Set //
// Set a value into the document tree
func (d *DocumentStore) Set(key uint32, value string, expiry time.Time) (COMPLETED, error) {

	err := d.Tree.SET(key, value, expiry)
	if err != nil {
		return COMPLETED{}, err
	}
	//
	//Clean exit
	return COMPLETED{
		//
		Key: key,
	}, nil
}

// // Update //
// // Update a value on the document tree
// func (d *DocumentStore) Update(key uint64, value string, expiry time.Time) (COMPLETED, error) {
// 	//
// 	return COMPLETED{
// 		//
// 		Key:    key,
// 		Ok:     true,
// 		AtTime: time.Now(),
// 	}, nil
// }

// Delete //
// Delete a value from the document tree
func (d *DocumentStore) Delete(key uint32) (COMPLETED, error) {
	err := d.Tree.DELETE(key)
	if err != nil {
		return COMPLETED{}, err
	}
	return COMPLETED{
		//
		Key:    key,
		AtTime: time.Now(),
	}, nil
}

// ############  sanic data structure

//SanicBST ..
type SanicBST struct {
	//Data storage
	tree treeStruct
	//interface
	sb
}
type sb interface {
	//Public
	TestTree()
	Init()
	//Public CMD
	SET(uint32, string, time.Time) error
	GET(uint32) (interface{}, error)
	UPDATE()
	DELETE(uint32) error
	//Node
	writeNewNode(uint32, string, time.Time) (*nodeStruct, error)
	nilNodeandQueue(n *nodeStruct) error
	//Tree
	insert(uint32, []byte, int64) error
	update(uint32, string, time.Time) error
	findNode(uint32) *nodeStruct
	findValue(uint32) []byte
	deleteNode(uint32) error
	//Recursive	functions
	findClosestR(uint32, *nodeStruct) *nodeStruct
	Inorder()
	find(*nodeStruct, uint32) (*nodeStruct, error)
}

//SET ..
func (s *SanicBST) SET(k uint32, v string, e time.Time) error {

	//todo
	//convert time.time to int64 somehow ----- RESEARCH <----
	// string to byte -- find fastest solution
	return s.insert(k, []byte(v), 0)

}

//GET ..
func (s *SanicBST) GET(k uint32) ([]byte, error) {
	//todo
	//convert time.time to int64 somehow ----- RESEARCH <----
	// string to byte -- find fastest solution

	nodeVal, err := s.find(k)
	if err != nil {
		return nil, err
	}
	val := nodeVal.value

	return val, nil

}

//DELETE ..
func (s *SanicBST) DELETE(k uint32) error {
	//todo
	//convert time.time to int64 somehow ----- RESEARCH <----
	// string to byte -- find fastest solution

	fmt.Println("Not yet impletemented")
	return nil

}
