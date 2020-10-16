package server

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/lafskelton/sanicdb/pkg/lexistore"
	"github.com/lafskelton/sanicdb/pkg/options"
	"github.com/lafskelton/sanicdb/pkg/store"
)

//Testing
const (
	testingEnabled = true
	testDocName    = "test"
)

//SanicDB Main data struct
type SanicDB struct {
	status  sanicStatus
	options options.SanicService
	docs    map[string]*store.DocumentStore
	ldocs   map[string]*lexistore.DocumentStore
	S
	Connected []net.Conn
}

//Allows components to pass errors up
type sanicStatus struct {
	ok  bool
	err error
}

//S public methods
type S interface {
	Test()
	Network()
	NewNumsDocumentStore(string)
	NewLexiDocumentStore(string)
	CreateNumsDoc(string, func(store.COMPLETED)) error
	CreateLexiDoc(string, func(store.COMPLETED)) error
	//numerical data
	Get(string, uint32, func(store.COMPLETED)) (string, error)
	Set(string, uint32, string, func(store.COMPLETED))
	Del(string, uint32, func(store.COMPLETED)) error
	//lexi data
	GetLexi(string, uint32, func(store.COMPLETED)) (string, error)
	SetLexi(string, uint32, string, func(store.COMPLETED))
	DelLexi(string, uint32, func(store.COMPLETED)) error
}

//StartService ..
//Returns an instance of sanicDB
func StartService() *SanicDB {
	opt := options.SanicService{
		Port: "42069",
	}
	sanic := SanicDB{
		options: opt,
		status: sanicStatus{
			ok:  true,
			err: nil,
		},
		docs:  make(map[string]*store.DocumentStore),
		ldocs: make(map[string]*lexistore.DocumentStore),
	}
	sanic.NewNumsDocumentStore("test")
	sanic.NewLexiDocumentStore("test")
	// this enables gRPC in a routine

	return &sanic
}

//#### Methods

//NewNumsDocumentStore allocs a new store and appends it to the db
func (s *SanicDB) NewNumsDocumentStore(name string) error {
	//Alloc document store
	if _, ok := s.docs[name]; ok {
		return errors.New("exists")
	}
	doc := new(store.DocumentStore)
	err := doc.Start(name)
	if err != nil {
		return errors.New("failedToStart")
	}
	//add to database
	s.docs[name] = doc
	return nil
}

//NewLexiDocumentStore allocs a new store and appends it to the db
func (s *SanicDB) NewLexiDocumentStore(name string) error {
	fmt.Println("Not implemented")
	return nil
}

//CreateNumsDoc ..
func (s *SanicDB) CreateNumsDoc(doc string, resp func(store.COMPLETED)) {
	err := s.NewNumsDocumentStore(doc)
	if err != nil {
		resp(store.COMPLETED{
			WithError: err,
		})
		return
	}
	var val interface{} = doc
	resp(store.COMPLETED{
		TaskType:  "doc",
		Key:       0,
		Value:     val,
		AtTime:    time.Now(),
		WithError: nil,
	})
	return
}

//CreateLexiDoc ..
func (s *SanicDB) CreateLexiDoc(doc string, resp func(store.COMPLETED)) {
	fmt.Println("Not implemented")
	return
}

//TODO DELETE METHOD

//Set ..  set
func (s *SanicDB) Set(doc string, key uint32, value string, resp func(interface{})) {
	exp := time.Now().Add(time.Minute)
	//Check if doc exists
	if _, ok := s.docs[doc]; ok {
		//do something here
		//Send set task to db
		task := store.SET{
			Resp:   resp,
			Key:    key,
			Value:  value,
			Expiry: exp,
		}
		s.docs[doc].Pool.Task.Queue.New(task)
		s.docs[doc].Pool.Task.NewTask.Signal()
		return
	}
	resp(store.COMPLETED{
		TaskType:  "set",
		Key:       0,
		Value:     nil,
		AtTime:    time.Now(),
		WithError: errors.New("noDoc"),
	})

	return
}

//Get ..
func (s *SanicDB) Get(doc string, key uint32, resp func(interface{}), wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	//Check if doc exists
	if _, ok := s.docs[doc]; ok {
		//do something here
		//Send set task to db
		str, err := s.docs[doc].Get(key)
		if err != nil {
			resp(store.COMPLETED{
				WithError: err,
			})
			return
		}
		resp(store.COMPLETED{
			TaskType:  "get",
			Key:       key,
			Value:     str,
			AtTime:    time.Now(),
			WithError: nil,
		})
		return
	}
	resp(store.COMPLETED{
		TaskType:  "get",
		Key:       0,
		Value:     nil,
		AtTime:    time.Now(),
		WithError: errors.New("noDoc"),
	})
	return
}

//Del ..
func (s *SanicDB) Del(doc string, key uint32, resp func(interface{})) error {
	//Check if doc exists
	if _, ok := s.docs[doc]; ok {
		//do something here
		//Send set task to db
		task := store.DELETE{
			Resp: resp,
			Key:  key,
		}

		s.docs[doc].Pool.Task.Queue.New(task)
		s.docs[doc].Pool.Task.NewTask.Signal()
		return nil
	}
	return errors.New("noDoc")
}

//Lexi Data

//SetLexi ..  set
func (s *SanicDB) SetLexi(doc string, key string, value string, resp func(interface{})) {
	fmt.Println("Not implemented")
	resp(nil)
	return
}

//GetLexi ..
func (s *SanicDB) GetLexi(doc string, key string, resp func(interface{})) {
	fmt.Println("Not implemented")
	resp(nil)
	return
}
