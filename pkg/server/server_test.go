package server

import (
	"context"
	"fmt"
	pb "github.com/bookish-goggles/protogen"
	"google.golang.org/grpc"
	"testing"
	"time"
	"math/rand"
)

// These tests assume the kvstore server is running

// Initial setup to establish connection with server
func connectToKVStore() (context.Context, context.CancelFunc, pb.KVStoreClient, *grpc.ClientConn, error) {
	// setup connection to oplang server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// caller should defer cancel func.
	client := pb.NewKVStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	return ctx, cancel, client, conn, nil
}

func TestGetAndPut(t *testing.T) {
	ctx, cancel, client, conn, err := connectToKVStore()
	if err != nil {
		t.Fatalf("Issue connecting to kvstore: %s", err)
	}
	defer cancel()
	defer conn.Close()
	putReq := &pb.PutReq{Key: "k", Val: "v"}
	_, err = client.Put(ctx, putReq)
	if err != nil {
		t.Fatalf("Issue executing put: %s", err)
	}
	getReq := &pb.GetReq{Key: putReq.Key}
	getRes, err := client.Get(ctx, getReq)
	if err != nil {
		t.Fatalf("Issue executing get: %s", err)
	}
	if getRes.Val != putReq.Val {
		t.Fatalf("Expected %s from get, but got %s", getRes.Val, putReq.Val)
	}
}

func BenchmarkPutSequential(b *testing.B) {
	ctx, cancel, client, conn, err := connectToKVStore()
	if err != nil {
		b.Fatalf("Issue connecting to kvstore: %s", err)
	}
	defer cancel()
	defer conn.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		putReq := &pb.PutReq{Key: fmt.Sprintf("k%d", n), Val: fmt.Sprintf("v%d", n)}
		_, err = client.Put(ctx, putReq)
		if err != nil {
			b.Fatalf("Issue executing put: %s", err)
		}
	}
}

func BenchmarkGetSequential(b *testing.B) {
	ctx, cancel, client, conn, err := connectToKVStore()
	if err != nil {
		b.Fatalf("Issue connecting to kvstore: %s", err)
	}
	defer cancel()
	defer conn.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		putReq := &pb.PutReq{Key: fmt.Sprintf("k%d", n), Val: fmt.Sprintf("v%d", n)}
		_, err = client.Put(ctx, putReq)
		if err != nil {
			b.Fatalf("Issue executing put: %s", err)
		}
	}
}

