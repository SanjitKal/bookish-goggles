package memtable

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type rbPathEntry struct {
	node *RbNode
	blackLength int //excludes the color of the stored node
}

func getKeyStr(node *RbNode) string {
	if node == nil {
		return "nil"
	} else if node.key == "" {
		return "sentinel"
	} else {
		return node.key
	}
}

func printRbNode(t *testing.T, node *RbNode) {
	if node == nil {
		return
	}

	colorStr := "black"
	if !node.color {
		colorStr = "red"
	}
	t.Logf("KEY: %s, COLOR: %s, PARENTKEY: %s, LEFTKEY: %s, RIGHTKEY: %s", getKeyStr(node), colorStr, getKeyStr(node.p), getKeyStr(node.l), getKeyStr(node.r))
}

func checkBstProperties(t *testing.T, memt *Memtable) {
	checkBstPropertiesHelper(t, memt.root)
}

func checkBstPropertiesHelper(t *testing.T, node *RbNode) (treeMax string, treeMin string) {
	if node == nil || node.key == "" {
		return "", ""
	}
	leftSubtreeMin, leftSubtreeMax := checkBstPropertiesHelper(t, node.l)
	rightSubtreeMin, rightSubtreeMax := checkBstPropertiesHelper(t, node.r)
	if leftSubtreeMax != "" && node.key < leftSubtreeMax {
		t.Fatalf("Found node with key=%s with a left subtree that contains a greater key=%s", node.key, leftSubtreeMax)
	} else if rightSubtreeMin != "" && node.key > rightSubtreeMin {
		t.Fatalf("Found node with key=%s with a right subtree that contains a lesser key=%s", node.key, rightSubtreeMin)		
	}
	treeMax = rightSubtreeMax
	treeMin = leftSubtreeMin
	if treeMin == "" {
		treeMin = node.key
	}
	if treeMax == "" {
		treeMax = node.key
	}
	return 
}

func checkRbProperties(t *testing.T, memt *Memtable) {
	// Check that the root is black
	if !memt.root.color {
		t.Fatal("Found red root")
	}

	// Perform BFS from root to check other red black properties
	queue := make([]rbPathEntry, 0)
	queue = append(queue, rbPathEntry{node: memt.root})
	allPathsBlackLength := -1
	var lastSentinelNodesParent *RbNode
	var curr rbPathEntry
	seen := make(map[*RbNode]bool)
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		printRbNode(t, curr.node)
		// Check red parent -> black child property
		if curr.node.p != nil && !curr.node.p.color && !curr.node.color {
			t.Fatalf("Found red parent=%s with red child=%s", curr.node.p.key, curr.node.key)
		}
		if curr.node.l == nil && curr.node.r == nil {
			// Ensure sentinel nodes are black
			if !curr.node.color {
				t.Fatalf("Found a red sentinel node with parent=%s", curr.node.p.key)
			}
			// Check black length property
			if allPathsBlackLength == -1 {
				allPathsBlackLength = curr.blackLength
			} else if allPathsBlackLength != curr.blackLength {
				t.Fatalf("Found path from root=%s to sentinel with parent=%s of black length=%d, but preceding root to sentinel with parent=%s had black length=%d", memt.root.key, curr.node.p.key, curr.blackLength, lastSentinelNodesParent.key, allPathsBlackLength)
			}
			lastSentinelNodesParent = curr.node.p
		} else {
			if seen[curr.node] {
				t.Fatalf("Encounted node=%s twice", curr.node.key)
			}
			seen[curr.node] = true
			// Add new path entries to queue
			newBlackLength := curr.blackLength
			if curr.node.color {
				newBlackLength += 1
			}
			if curr.node.l != nil {
				queue = append(queue, rbPathEntry{node: curr.node.l, blackLength: newBlackLength})
			}
			if curr.node.r != nil {
				queue = append(queue, rbPathEntry{node: curr.node.r, blackLength: newBlackLength})
			}
		}
	}
}

func lookupMemtableRange(numElts int, memt *Memtable) {
	for i := 0; i < numElts; i++ {
		memt.Lookup(fmt.Sprintf("k%d", i))
	}
}

func fillMemtableSeq(numElts int, memt *Memtable) (err error, keyArr []string, valArr []string) {
	keyArr = make([]string, numElts)
	valArr = make([]string, numElts)
	for i := 0; i < numElts; i++ {
		keyArr[i] = fmt.Sprintf("k%d", i)
		valArr[i] = fmt.Sprintf("v%d", i)
		err = memt.Insert(keyArr[i], valArr[i])
		if err != nil {
			return err, keyArr, valArr
		}
	}
	return nil, keyArr, valArr
}

func fillMemtableRandom(numElts int, memt *Memtable) (err error, keyArr []string, valArr []string) {
	keyArr = make([]string, numElts)
	valArr = make([]string, numElts)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numElts; i++ {
		randInt := rand.Intn(numElts)
		keyArr[i] = fmt.Sprintf("k%d", randInt)
		valArr[i] = fmt.Sprintf("v%d", randInt)
		err = memt.Insert(keyArr[i], valArr[i])
		if err != nil {
			return err, keyArr, valArr
		}
	}
	return nil, keyArr, valArr
}

func TestInsertSeqAndLookup(t *testing.T) {
	memt := new(Memtable)
	memt.Init(-1)
	err, keyArr, valArr := fillMemtableSeq(3, memt)
	if err != nil {
		t.Fatalf(err.Error())
	}
	checkBstProperties(t, memt)
	checkRbProperties(t, memt)
	for i := 0; i < len(keyArr); i++ {
		val, err := memt.Lookup(keyArr[i])
		if err != nil {
			t.Fatalf(err.Error())
		}
		if val != valArr[i] {
			t.Fatalf("Expected to find val=%s associated with key=%s, instead found val=%s", valArr[i], keyArr[i], val)
		}
	}
}

func TestInsertRandomAndLookup(t *testing.T) {
	memt := new(Memtable)
	memt.Init(-1)
	err, keyArr, valArr := fillMemtableRandom(1000000, memt)
	t.Log("insertion sequence:", keyArr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	checkBstProperties(t, memt)
	checkRbProperties(t, memt)
	for i := 0; i < len(keyArr); i++ {
		val, err := memt.Lookup(keyArr[i])
		if err != nil {
			t.Fatalf(err.Error())
		}
		if val != valArr[i] {
			t.Fatalf("Expected to find val=%s associated with key=%s, instead found val=%s", valArr[i], keyArr[i], val)
		}
	}
}

func TestInsertAtCapacity(t *testing.T) {
	memt := new(Memtable)
	memt.Init(9)
	err, _, _ := fillMemtableSeq(10, memt)
	if err == nil {
		t.Fatalf("Expected error due to insertion at capacity, but received none")
	}
}

func Benchmark1MilRandomInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		memt := new(Memtable)
		memt.Init(-1)
		fillMemtableRandom(1000000, memt)
	}
}

func Benchmark1MilSeqInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		memt := new(Memtable)
		memt.Init(-1)
		fillMemtableSeq(1000000, memt)
	}
}

func BenchmarkSeqLookupAllKeysAfter1MilSeqInsert(b *testing.B) {
	memt := new(Memtable)
	memt.Init(-1)
	fillMemtableRandom(1000000, memt)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lookupMemtableRange(1000000, memt)
	}
}

func BenchmarkSeqLookupAllKeysAfter1MilRandomInsert(b *testing.B) {
	memt := new(Memtable)
	memt.Init(-1)
	fillMemtableRandom(1000000, memt)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lookupMemtableRange(1000000, memt)
	}
}
