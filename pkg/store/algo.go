package store

import (
	"errors"
	"fmt"
	"math"
)

//This function recursively finds the closest node to the key in time O(log n)
func (s *SanicBST) findClosestR(key uint32) (*nodeStruct, error) {
	var curs *nodeStruct = s.tree.root
	var minDiff int64
	var minPtr *nodeStruct
	for {
		//Closest
		diff := int64(math.Abs(float64((int64(key) - int64(curs.key)))))
		if diff < minDiff {
			minDiff = diff
			minPtr = curs
		}
		//If key is less than cursor
		if key < curs.key {
			//If the
			if curs.left != nil {
				curs = curs.left
				continue
			}
		}

		//If key is greater than cursor
		if key > curs.key {
			if curs.right != nil {
				curs = curs.right
				continue
			}
		}

		//If node is key
		if curs.key == key && curs.key != s.tree.root.key {
			return curs, nil
		} else if curs.key == 0 {
			return curs, errors.New("noNode " + fmt.Sprint(key))
		}
		//Return Node

		if minPtr != nil {
			return minPtr, nil
		}
		return curs, nil
	}

}

//Recursive function that returns the minimum node below the cursor
func (s *SanicBST) findMin(cursor *nodeStruct) (*nodeStruct, error) {
	var curs *nodeStruct = cursor
	if cursor == nil {
		curs = s.tree.root
	}
	for curs.left != nil {
		curs = curs.left
	}
	fmt.Println(curs)
	return curs, nil
}

//Recursive function returns a node matching key
func (s *SanicBST) find(k uint32) (*nodeStruct, error) {
	var curs *nodeStruct
	curs = s.tree.root
	for {
		//If key is less than cursor
		if k < curs.key {
			//If the
			if curs.left != nil {
				curs = curs.left
				continue
			}
		}
		//If key is greater than cursor
		if k > curs.key {
			if curs.right != nil {
				curs = curs.right
				continue
			}
		}
		//If node is key
		if curs.key == k && curs.key != s.tree.root.key {
			return curs, nil
		} else if curs.key == 0 {
			return curs, errors.New("noNode " + fmt.Sprint(k))
		}
		return nil, errors.New("notFound2")
	}
}

//Inorder prints out the tree in order
func (s *SanicBST) Inorder(cursor *nodeStruct) {
	//If starting point is nil, start at root
	if cursor == nil {
		cursor = s.tree.root
	}
	if cursor.left != nil {
		s.Inorder(cursor.left)
	}
	fmt.Println(*cursor)
	if cursor.right != nil {
		s.Inorder(cursor.right)
	}
	return
}
