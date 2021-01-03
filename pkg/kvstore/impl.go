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
}

func (kv *KVStore) Get(key string) (string, pb.Error) {
	val, err := kv.memt.Lookup(key)
	if err != nil {
		return "", pb.Error{Type: pb.Error_GET_ERROR, Message: fmt.Sprintf(`key "%s" not found`, key)}
	}
	return val, pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) Put(key string, val string) pb.Error {
	err := kv.memt.Insert(key, val)
	if err != nil {
		return pb.Error{Type: pb.Error_PUT_ERROR, Message: err.Error()}
	}
	return pb.Error{Type: pb.Error_NO_ERROR}

}

func (kv *KVStore) Del(key string) pb.Error {
	return pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) Load() pb.Error {
	return pb.Error{Type: pb.Error_NO_ERROR}
}
