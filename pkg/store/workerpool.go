package store

import (
	"sync"
)

//WorkerPool handles tasks to be executed on the heap
type WorkerPool struct {
	//This workers store
	store *DocumentStore
	//Task Pool
	Task     TaskPool
	Finished Complete
	//Worker function
	Workers
	//Error queue
	Errs errorQueue
	//Tree Collapse Callback
	collapse func() error
}

//TaskPool ...
type TaskPool struct {
	//Mutex
	Mu sync.Mutex
	//Queue of waiting tasks
	Queue taskQueue
	//NewTask signals the task worker to work
	NewTask sync.Cond
}

//Complete //
type Complete struct {
	//Mutex
	Mu          sync.Mutex
	NewComplete sync.Cond
	//Finished tasks ready to send responce
	Queue compQueue
}

//Workers .. Worker Function
type Workers interface {
	StartAll()
	//This waits for tasks and executes them respectively
	TaskWorker()
	//This signals a task has finished
	finishTask(COMPLETED)
	//MemoryWorker on a set interval, the memory worker collaspes the forest
	MemoryWorker()
	//ExpiryWorker
	ExpiryWorker()
}

//Worker pool methods

//StartAll ..
func (w *WorkerPool) StartAll() {
	go w.TaskWorker()
	go w.ExpiryWorker()
}

//TaskWorker ..
func (w *WorkerPool) TaskWorker() {
	for {
		w.Task.Mu.Lock()
		//Block & wait if no jobs
		for w.Task.Queue.list.Len() == 0 {
			//No, Wait for jobs
			w.Task.NewTask.Wait()
		}
		//Pop
		j := w.Task.Queue.Pop()
		//Type switch and execute task
		switch task := j.(type) {
		case SET:
			completed, err := w.store.Set(task.Key, task.Value, task.Expiry)
			if err != nil {

				//###############
				// Change this to invoke the gRPC Set callback with an error instead of
				// using an error task worker
				//////

				completed.TaskType = "SET"
				completed.Doc = w.store.Name
				completed.Resp = task.Resp
				completed.Key = task.Key
				completed.WithError = err
				w.finishTask(completed)
				//Continue
				w.Task.Mu.Unlock()
				continue
			}
			completed.TaskType = "SET"
			completed.Doc = w.store.Name
			completed.Resp = task.Resp
			completed.Key = task.Key
			w.finishTask(completed)
			//Continue
			w.Task.Mu.Unlock()
			continue
		case DELETE:
			completed, err := w.store.Delete(task.Key)
			if err != nil {
				//Fail task
				completed.TaskType = "DELETE"
				completed.Resp = task.Resp
				completed.Doc = w.store.Name
				completed.Key = task.Key
				completed.WithError = err
				w.finishTask(completed)
				//Continue
				w.Task.Mu.Unlock()
			}
			completed.TaskType = "DELETE"
			completed.Resp = task.Resp
			completed.Doc = w.store.Name
			completed.Key = task.Key
			w.finishTask(completed)
			//Continue
			w.Task.Mu.Unlock()
			continue
		}
		w.Task.Mu.Unlock()
	}
}

//This adds the task to the finished queue and notifies the completed queue
func (w *WorkerPool) finishTask(task COMPLETED) {
	task.TaskType = "pool"
	task.Resp(task)
	return

}

//ExpiryWorker ..
func (w *WorkerPool) ExpiryWorker() {
	// for {
	// 	//Make things expire
	// }
}
