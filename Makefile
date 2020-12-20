all: kvstore cli
	gofmt -w .

kvstore: gen-protos
	go build -o bin/kvstore ./cmd/kvstore

gen-protos:
	./scripts/gen_protos.sh

cli: gen-protos
	go build -o bin/cli ./cmd/cli

