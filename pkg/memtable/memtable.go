package memtable

// Implements memtable interface, which uses a red black tree as its' underlying data structure. Therefore, insertion of keys
// into the memtable is O(logn), and retrieving/removing the contents of the memtable in sorted order is O(nlogn).

import (
	"errors"
	"fmt"
)

type Memtable struct {
	capacity int     // unbounded capacity is represented by -1
	size     int     // number of keys stored in tree
	root     *RbNode // root of rb tree
}

type RbNode struct {
	p     *RbNode // parent
	l     *RbNode // left child
	r     *RbNode // right child
	color bool    // true if black, false if red
	key   string
	val   string
}

func (memt *Memtable) Init(capacity int) {
	memt.capacity = capacity
}

func (memt *Memtable) Lookup(key string) (val string, err error) {
	// Standard BST read
	curr := memt.root
	for curr != nil {
		if key == curr.key {
			return curr.val, nil
		} else if key < curr.key {
			curr = curr.l
		} else {
			curr = curr.r
		}
	}
	return "", errors.New(fmt.Sprintf(`key, "%s" ,not found in memtable`, key))
}

func (memt *Memtable) Insert(key string, val string) error {
	// Standard BST Insertion followed by restoration of RB properties

	if memt.root == nil {
		// Empty tree case
		memt.root = &RbNode{color: true, key: key, val: val}
		return nil
	}

	newNode := &RbNode{color: false, key: key, val: val}

	curr := memt.root
	for curr != nil {
		if key == curr.key {
			// Overwrite the value of the key if it already exists in the tree
			// Bypasses capacity check at end of function because we overwrite existing node.
			curr.val = val
			return nil
		} else if key < curr.key {
			if curr.l == nil {
				curr.l = newNode
				newNode.p = curr
			} else {
				curr = curr.l
			}
		} else {
			if curr.r == nil {
				curr.r = newNode
				newNode.p = curr
			} else {
				curr = curr.r
			}
		}
	}

	if memt.size == memt.capacity {
		return errors.New(fmt.Sprintf("Cannot insert (key=%s, val=%s), because at capacity=%d", key, val, memt.capacity))
	}

	memt.insertionFixup(newNode)
	memt.size++
	return nil
}

func (memt *Memtable) leftRotate(curr *RbNode) {
	// Left rotate curr by shifting it down and left,
	// and bringing it's right child to its' place.
	if curr.r == nil || curr.l == nil {
		return
	}

	curr.r.p = curr.p // curr's right child's parent is now curr's parent
	curr.p = curr.r   // curr's parent is now curr's right child
	curr.r = curr.r.l // curr's right child is now it's original right child's left subtree
	curr.r.l = curr   // curr's right child's left child now points to curr
}

func (memt *Memtable) rightRotate(curr *RbNode) {
	// Right rotate curr by shifting it down and right,
	// and bringing it's left child to its' place.
	// Symmetrical to leftRotate with left and right swapped.
	if curr.r == nil || curr.l == nil {
		return
	}

	curr.l.p = curr.p
	curr.p = curr.l
	curr.l = curr.l.r
	curr.l.r = curr
}

func (memt *Memtable) insertionFixup(curr *RbNode) {
	// Restore RB properties after insertion via repeated recolorings and rotations
	// until we reach the root of the tree, or the parent is black
	if curr == nil {
		return
	}

	for curr.p.color && curr.p != nil {
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
			// to the corresponding if stmt above, with left and right reversed.
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

func (memt *Memtable) Delete(key string) (err error) {
	return nil
}

func (memt *Memtable) PopMin() (key string, val string, err error) {
	return "", "", nil
}

func (memt *Memtable) PopAll() (key []string, val []string, err error) {
	return nil, nil, nil
}

func (memt *Memtable) GetSize() (size int) {
	return memt.size
}

func (memt *Memtable) GetCapacity() (cap int) {
	return memt.capacity
}
