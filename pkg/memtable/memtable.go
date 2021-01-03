package memtable

// Implements memtable interface, which uses a red black tree as its' underlying data structure. Therefore, insertion of keys
// into the memtable is O(logn), and retrieving/removing the contents of the memtable in sorted order is O(nlogn).

import (
	"errors"
	"fmt"
	"sync"
)

type Memtable struct {
	capacity int        // unbounded capacity is represented by -1. otherwise, capacity must be positive
	size     int        // number of keys stored in tree
	root     *RbNode    // root of rb tree
	sentinel *RbNode    // a keyless black node, which is used to represent leaves in the rb tree
	rwLock   sync.RWMutex // coarse grained reader writer lock to allow for multiple concurrent readers and a single writer
}

type RbNode struct {
	p     *RbNode // parent
	l     *RbNode // left child
	r     *RbNode // right child
	color bool    // true if black, false if red
	key   string  // rb node order is established using the key
	val   string  //
}

func (memt *Memtable) Init(capacity int) (err error) {
	err = memt.UpdateCapacity(capacity)
	if err != nil {
		return err
	}
	memt.sentinel = &RbNode{color: true} // sentinel node is always colored black
	memt.root = memt.sentinel
	return nil
}

func (memt *Memtable) UpdateCapacity(capacity int) error {
	memt.rwLock.Lock()
	defer memt.rwLock.Unlock()

	if capacity < 0 && capacity != -1 {
		return errors.New("Capacity must either be a positive integer, or -1 to represent an unbounded capacity.")
	}
	memt.capacity = capacity
	return nil
}

func (memt *Memtable) Lookup(key string) (val string, err error) {
	// Standard BST read
	memt.rwLock.RLock()
	defer memt.rwLock.RUnlock()

	curr := memt.root

	for curr != memt.sentinel {
		if key == curr.key {
			return curr.val, nil
		} else if key < curr.key {
			curr = curr.l
		} else {
			curr = curr.r
		}
	}
	return "", errors.New(fmt.Sprintf(`key, "%s",not found in memtable`, key))
}

func (memt *Memtable) Insert(key string, val string) error {
	// Standard BST Insertion followed by restoration of RB properties
	memt.rwLock.Lock()
	defer memt.rwLock.Unlock()

	newNode := &RbNode{color: false, key: key, val: val, l: memt.sentinel, r: memt.sentinel}

	if memt.root == memt.sentinel {
		// Empty tree case
		memt.root = newNode
		memt.root.color = true // set root to black
		memt.size++
		return nil
	}

	curr := memt.root
	for curr != memt.sentinel {
		if key == curr.key {
			// Overwrite the value of the key if it already exists in the tree
			curr.val = val
			return nil
		} else if key < curr.key {
			if curr.l == memt.sentinel {
				// Insert as left child
				curr.l = newNode
				newNode.p = curr
				break
			} else {
				curr = curr.l
			}
		} else {
			if curr.r == memt.sentinel {
				// Insert as right child
				curr.r = newNode
				newNode.p = curr
				break
			} else {
				curr = curr.r
			}
		}
	}

	if memt.size == memt.capacity {
		return errors.New(fmt.Sprintf(`Cannot insert new key="%s" because at capacity=%d`, key, memt.capacity))
	}

	memt.size++
	memt.insertionFixup(newNode)
	return nil
}

func (memt *Memtable) leftRotate(curr *RbNode) {
	// Left rotate curr by shifting it down and left,
	// and bringing it's right child to its' place.
	if curr.r == memt.sentinel {
		// sentinel nodes cannot become internal nodes
		return
	}

	if curr.p == nil {
		// reassign root
		memt.root = curr.r
	} else {
		// adjust the appropriate child pointer
		// of curr's parent
		if curr.p.l == curr {
			curr.p.l = curr.r
		} else {
			curr.p.r = curr.r
		}
	}

	// shift curr and it's right child
	originalRight := curr.r
	originalRightsLeftChild := curr.r.l
	originalRight.p = curr.p
	curr.p = originalRight
	curr.r.l = curr
	curr.r = originalRightsLeftChild
	originalRightsLeftChild.p = curr
}

func (memt *Memtable) rightRotate(curr *RbNode) {
	// Symmetric to leftRotate with left and right swapped.
	if curr.l == memt.sentinel {
		return
	}

	if curr.p == nil {
		memt.root = curr.l
	} else {
		if curr.p.r == curr {
			curr.p.r = curr.l
		} else {
			curr.p.l = curr.l
		}
	}

	originalLeft := curr.l
	originalLeftsRightChild := curr.l.r
	originalLeft.p = curr.p
	curr.p = originalLeft
	originalLeft.r = curr
	curr.l = originalLeftsRightChild
	originalLeftsRightChild.p = curr
}

func (memt *Memtable) insertionFixup(curr *RbNode) {
	// Restore RB properties after insertion via repeated recolorings and rotations
	// until we reach the root of the tree, or the parent is black
	for curr.p != nil && !curr.p.color {
		// curr's parent is red, we need to restore [red node --> black child] property.
		// we know the parent is not the root, because the root is always black.
		if curr.p.p.l == curr.p {
			// parent is left child of grandparent. must case on uncle's color
			if !curr.p.p.r.color {
				// uncle is red --> color uncle and parent black, color grandparent red
				// and repeat with grandparent as curr (since its' parent maybe red)
				curr.p.p.r.color = true
				curr.p.color = true
				curr.p.p.color = false
				curr = curr.p.p
			} else {
				// uncle is black
				if curr.p.r == curr {
					// curr is parent's right child --> make curr point to curr's parent, and
					// perform left rotation on parent.
					parent := curr.p
					curr = curr.p
					memt.leftRotate(parent)
				}

				// We can be now be sure that:
				// (1) curr (red) is a left child of curr's parent (red),
				// (2) curr's parent is a left child of curr's grandparent (black)
				// So, we color curr's parent black, curr's grandparent red, and
				// perform right rotation on curr's grandparent to restore
				// [root-leaves black node] property
				curr.p.color = true
				curr.p.p.color = false
				memt.rightRotate(curr.p.p)
				// Last iteration of loop because curr's parent is now red
			}
		} else {
			// parent is right child of grandparent. This case is symmetrical
			// to the corresponding if stmt above, with left and right swapped.
			if !curr.p.p.l.color {
				curr.p.p.l.color = true
				curr.p.color = true
				curr.p.p.color = false
				curr = curr.p.p
			} else {
				if curr.p.l == curr {
					parent := curr.p
					curr = curr.p
					memt.rightRotate(parent)
				}
				curr.p.color = true
				curr.p.p.color = false
				memt.leftRotate(curr.p.p)
			}
		}
	}

	// Ensure the root is black
	memt.root.color = true
}

func (memt *Memtable) getSortedEntriesByKeyHelper(curr *RbNode, keyArr []string, valArr []string, idx int) (nextIdx int) {
	if curr == memt.sentinel {
		return idx
	}
	nextIdx = memt.getSortedEntriesByKeyHelper(curr.l, keyArr, valArr, idx)
	keyArr[nextIdx] = curr.key
	valArr[nextIdx] = curr.val
	nextIdx = memt.getSortedEntriesByKeyHelper(curr.r, keyArr, valArr, nextIdx+1)
	return nextIdx
}

func (memt *Memtable) GetSortedEntriesByKey() (keyArr []string, valArr []string) {
	// Returns sorted key entries (in lexicographically ascending order)
	// and corresponding value entries: (keyArr[i], valArr[i]) represents
	// a single kv pair.
	memt.rwLock.RLock()
	defer memt.rwLock.RUnlock()

	keyArr = make([]string, memt.GetSize())
	valArr = make([]string, memt.GetSize())
	memt.getSortedEntriesByKeyHelper(memt.root, keyArr, valArr, 0)
	return keyArr, valArr
}

func (memt *Memtable) Clear() {
	// Free the memory allocated to store the entries in the rb tree.
	memt.rwLock.Lock()
	defer memt.rwLock.Unlock()

	memt.size = 0
	memt.root = memt.sentinel
}

func (memt *Memtable) GetSize() (size int) {
	memt.rwLock.RLock()
	defer memt.rwLock.RUnlock()

	return memt.size
}

func (memt *Memtable) GetCapacity() (cap int) {
	memt.rwLock.RLock()
	defer memt.rwLock.RUnlock()

	return memt.capacity
}
