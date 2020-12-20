package kvstore

import (
	"fmt"
	pb "github.com/bookish-goggles/protogen"
)

type KVStore struct {
	Store map[string]string
}

func (kv *KVStore) Init() {
	kv.Store = make(map[string]string)
}

func (kv *KVStore) Get(key string) (string, pb.Error) {
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
