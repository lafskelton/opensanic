package store

import (
	"time"
)

//#### Numerical Task Queue Structs

//GET operation
type GET struct {
	Resp func(interface{}) //Will accept either lexistore.COMPLETED or store.COMPLETED
	Key  uint32
	err  error
}

//SET operation
type SET struct {
	Resp   func(interface{}) //Will accept either lexistore.COMPLETED or store.COMPLETED
	Key    uint32
	Value  string
	Expiry time.Time
	Done   chan struct{}
	err    error
}

//UPDATE operation
type UPDATE struct {
	Resp   func(interface{}) //Will accept either lexistore.COMPLETED or store.COMPLETED
	Key    uint32
	Value  string
	Expiry time.Time
	err    error
}

//DELETE operation
type DELETE struct {
	Resp func(interface{}) //Will accept either lexistore.COMPLETED or store.COMPLETED
	Key  uint32
	err  error
}

//COMPLETED //
// data to send to the requesting client
type COMPLETED struct {
	Resp      func(interface{}) //Will accept either lexistore.COMPLETED or store.COMPLETED
	TaskType  string
	Doc       string
	Key       uint32
	Value     interface{}
	AtTime    time.Time
	WithError error
}
