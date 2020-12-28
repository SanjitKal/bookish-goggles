package kvstore

import (
	"fmt"
	pb "github.com/bookish-goggles/protogen"
	mem "github.com/bookish-goggles/pkg/memtable"
	lsm "github.com/bookish-goggles/pkg/lsm"
	wal	"github.com/bookish-goggles/pkg/wal"
)

type KVStore struct {
	Store map[string]string
	MemTable mem.Memtable
	Lsm lsm.LogStructuredMergeTree
	Wal wal.WriteAheadLog
}

func (kv *KVStore) Init() {
	kv.Store = make(map[string]string)
	// TODO:
	// initialize memtable instance
	// initialize lsm instance
	// initialize wal instance
	// call load, which runs operations from wal
}

func (kv *KVStore) Get(key string) (string, pb.Error) {
	// TODO:
	// write to wal
	// write to memtable
	// if memtable capacity reached:
		// write contents of memtable to lsm
		// clear memtable
		// clear wal
	if val, ok := kv.Store[key]; ok {
		return val, pb.Error{Type: pb.Error_NO_ERROR}
	} else {
		msg := fmt.Sprintf(`key "%s" not found in kvstore`, key)
		return "", pb.Error{Type: pb.Error_KEY_NOT_FOUND, Message: msg}
	}
}

func (kv *KVStore) Put(key string, val string) pb.Error {
	kv.Store[key] = val
	return pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) Del(key string) pb.Error {
	return pb.Error{Type: pb.Error_NO_ERROR}
}

func (kv *KVStore) Load() pb.Error {
	return pb.Error{Type: pb.Error_NO_ERROR}
}

