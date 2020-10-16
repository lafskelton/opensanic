package store

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

const (
	branchSize = 11111
)

type treeStruct struct {
	//Locker
	sync.Mutex
	//Indexes
	nextNil       uint
	currentBranch uint32
	root          *nodeStruct
	//Data
	stack []branch
	//Queue
	empty *emptyQueue
	ts
}
type ts interface {
	newBranch()
}

func (ts *treeStruct) newBranch() {
	//New branch
	ts.stack = append(ts.stack, branch{
		nodes: [branchSize]nodeStruct{},
	})
	//public indexes
	ts.currentBranch++
	ts.nextNil = 0
	return
}

type branch struct {
	nodes [branchSize]nodeStruct
}

//NodeStruct struct //
type nodeStruct struct {
	//Locker
	live        bool //indicates if the node is live or not
	key         uint32
	value       []byte
	expiry      int64 // For Expiry worker
	left        *nodeStruct
	right       *nodeStruct
	parent      *nodeStruct //Parent tracking allows the delete algo to stitch the tree together
	treeIndex   uint32
	branchIndex uint
}

//PUBLIC

//Init .. initializes SanicBST tree
func (s *SanicBST) Init() error {
	//Initialize Queues
	s.tree = treeStruct{
		stack: []branch{{
			nodes: [branchSize]nodeStruct{}},
		},
		empty:         new(emptyQueue),
		nextNil:       0,
		currentBranch: 0,
	}
	s.tree.empty.list = list.New()
	s.tree.root = s.writeNewNode(uint32(branchSize/2), []byte("root"), int64(0))
	return nil
}

//Node Methods

//writeNewNode ..
func (s *SanicBST) writeNewNode(key uint32, value []byte, exp int64) *nodeStruct {
	var n *nodeStruct
	var i uint
	//Check if there is an empty node, if not, look at the branch and create a new one
	if empty := s.tree.empty.Pop(); empty != nil {
		n = empty
	} else {
		//Check if the branch is full
		if s.tree.nextNil >= branchSize-1 {
			s.tree.newBranch()
		}
		//Write to the empty node
		i = s.tree.nextNil
		s.tree.nextNil++

		n = &s.tree.stack[s.tree.currentBranch].nodes[i]
		n.treeIndex = s.tree.currentBranch
		n.branchIndex = i
	}

	n.live = true
	n.key = key
	n.value = value
	n.expiry = exp

	//Inc. write index
	//Return
	return n
}

func (s *SanicBST) nilNodeandQueue(n *nodeStruct) {

	n.live = false
	n.key = 0
	n.value = nil
	n.expiry = 0

	n.parent = nil
	n.left = nil
	n.right = nil

	//Add to empty queue
	s.tree.empty.Push(n)

	return
}

//Tree Methods

//insert a new node into the tree
func (s *SanicBST) insert(key uint32, value []byte, exp int64) error {
	//Writes start here
	s.tree.Lock()
	defer s.tree.Unlock()
	//Fork a process to alloc a new node and signal when ready
	n := s.writeNewNode(key, value, exp)
	closest, err := s.findClosestR(key)
	if err != nil {
		return err
	}
	// The closest has the same key, this should only be done via update return an error
	if key == closest.key {
		if key == s.tree.root.key {
			s.tree.root.value = value
			return nil
		}
		//Add node to deleted queue
		s.nilNodeandQueue(n)
		return errors.New("nodeExists" + fmt.Sprint(key))
	}
	//Cases

	// closest is a leaf, set as left
	if closest.left == nil && closest.right == nil {

		if key < closest.key {
			closest.left = n
			n.parent = closest
		}
		if key > closest.key {
			closest.right = n
			n.parent = closest
		}

		return nil
	}
	// Equal Key above
	// The key is less than the closest node
	// The key is greater than the closest node
	if key < closest.key {
		//If the left child is nil, this belongs there
		if closest.left == nil {
			n.parent = closest
			//Set closest as left child
			closest.left = n
			return nil
			//There is a left node, and the key is greater than it, is there a nil right node?
		} else if closest.right == nil {

			//Check that the new node is greater than the left node
			l := closest.left
			if key > l.key && l != nil {

				closest.right = n
				n.parent = closest
				return nil
			}
		}

		//move the left child down and put new node here
		//Set parent as closests parent
		n.parent = closest.parent
		closest.parent = n
		n.left = closest

		//the key is less than the closest

		//if closests right child is greater than new node, move it to new nodes right child
		r := closest.right
		if r != nil && key < r.key {
			//Set new nodes right child
			n.right = r
			//Update the moved nodes parent
			r.parent = n
			//Update the closest nodes right child
			closest.right = nil
		}
		return nil
	} else if key > closest.key {
		//If the right child is nil, this belongs there
		if closest.right == nil {

			n.parent = closest
			closest.right = n

			return nil
		}
		//move the left child down and put new node here
		//Set parent as closests parent

		n.parent = closest.parent
		closest.parent = n
		n.left = closest

		//if closests right child is greater than new node, move it to new nodes right child
		r := closest.right
		if key < r.key {
			//Set new nodes right child
			n.right = closest.right
			//Update the moved nodes parent
			r.parent = n
			//Update the closest nodes right child
			closest.right = nil
		}

		return nil
	}
	return nil
}
