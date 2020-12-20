package server

import (
	kvstore "github.com/bookish-goggles/pkg/kvstore"
)

var state *ServerState

type ServerState struct {
	KVStoreInstance *kvstore.KVStore
}

func (state *ServerState) Init() {
	state.KVStoreInstance = new(kvstore.KVStore)
	state.KVStoreInstance.Init()
}
