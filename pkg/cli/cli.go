// Simple cli to interact with a kvstore node over http.
package cli

import (
	"context"
	"fmt"
	pb "github.com/bookish-goggles/protogen"
	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"
)

const (
	timeout = time.Minute * 2
)

func StartCli() {
	app := cli.NewApp()
	app.Action = promptAction
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "get k", Description: `Retrieve the value of key 'k' from the store.`},
		{Text: "put k v", Description: "Insert/update the key 'k', bound to 'v'."},
		{Text: "exit", Description: "Exits the cli"},
	}
}

func handleGet(key string) (string, error) {
	ctx, cancel, client, conn, err := connectToKVStore()
	if err != nil {
		return "", err
	}
	defer cancel()
	defer conn.Close()
	getReq := &pb.GetReq{Key: key}
	getRes, err := client.Get(ctx, getReq)
	if err != nil {
		return "", err
	} else if getRes.Err.Type != pb.Error_NO_ERROR {
		return getRes.Err.Message, nil
	} else {
		return getRes.Val, nil
	}
}

func handlePut(key string, val string) (string, error) {
	ctx, cancel, client, conn, err := connectToKVStore()
	if err != nil {
		return "", err
	}
	defer cancel()
	defer conn.Close()
	putReq := &pb.PutReq{Key: key, Val: val}
	putRes, err := client.Put(ctx, putReq)
	if err != nil {
		return "", err
	} else if putRes.Err.Type != pb.Error_NO_ERROR {
		return putRes.Err.Message, nil
	} else {
		return "", nil
	}
}

func connectToKVStore() (context.Context, context.CancelFunc, pb.KVStoreClient, *grpc.ClientConn, error) {
	// Connect to kvstore
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client := pb.NewKVStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	return ctx, cancel, client, conn, nil
}

func executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	if input == "exit" {
		os.Exit(0)
		return
	}

	words := strings.Split(input, " ")

	switch words[0] {
	case "get":
		if len(words) != 2 {
			fmt.Println(`Invalid syntax. Example usage: "get somekey"`)
			return
		}
		res, err := handleGet(words[1])
		if err != nil {
			fmt.Println("Issue executing get:", err)
		} else {
			fmt.Println(res)
		}
	case "put":
		if len(words) != 3 {
			fmt.Println(`Invalid syntax. Example usage: "put somekey someval"`)
			return
		}
		res, err := handlePut(words[1], words[2])
		if err != nil {
			fmt.Println("Issue executing put:", err)
		} else if res != "" {
			fmt.Println(res)
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Println(`Unrecognized command. Supported commands: get, put, exit.`)
	}

	return
}

func promptAction(ctx *cli.Context) error {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("kvstore-cli$ "),
		prompt.OptionMaxSuggestion(3),
	)
	p.Run()
	return nil
}
