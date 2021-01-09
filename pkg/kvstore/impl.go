package kvstore

import (
	"fmt"
	lsm "github.com/bookish-goggles/pkg/lsm"
	mem "github.com/bookish-goggles/pkg/memtable"
	wal "github.com/bookish-goggles/pkg/wal"
	pb "github.com/bookish-goggles/protogen"
)

type KVStore struct {
	memt *mem.Memtable
	lsm  *lsm.LogStructuredMergeTree
	wal  *wal.WriteAheadLog
}

func (kv *KVStore) Init() {
	kv.memt = &mem.Memtable{}
	kv.memt.Init(-1)
	kv.lsm.Init(-1)
}

func (kv *KVStore) Get(key string) (val string, pbErr pb.Error) {
	// Search in memory memtable first
	val, err := kv.memt.Lookup(key)
	if err == nil {
		return val, pb.Error{Type: pb.Error_NO_ERROR}
	}
	// Search on disk via lsm
	val, err = kv.lsm.Lookup(key)
	if err == nil {
		return val, pb.Error{Type: pb.Error_NO_ERROR}
	}
	return "", pb.Error{Type: pb.Error_GET_ERROR, Message: fmt.Sprintf(`key "%s" not found`, key)}
}

func (kv *KVStore) Put(key string, val string) (pbErr pb.Error) {
	// If memtable is full, flush
	if kv.memt.GetSize() == kv.memt.GetCapacity() {
		err := kv.flush()
		if err != nil {
			return pb.Error{Type: pb.Error_PUT_ERROR, Message: err.Error()}
		}
	}
	// Insert into memtable
	err := kv.memt.Insert(key, val)
	if err != nil {
		return pb.Error{Type: pb.Error_PUT_ERROR, Message: err.Error()}
	}
	return pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) Del(key string) (pbErr pb.Error) {
	// Internally, deleting a key is representing
	// by inserting the pair (key, "")
	kv.Put(key, "")
	return pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) load() (err error) {
	return nil
}

func (kv *KVStore) flush() (err error) {
	// Write the sorted memtable entries into the lsm
	// and clear the memtable
	sortedKeyArr, valArr := kv.memt.GetSortedEntriesByKey()
	kv.memt.Clear()
	return kv.lsm.WriteNewSSTable(sortedKeyArr, valArr)
}
