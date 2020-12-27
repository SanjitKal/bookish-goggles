all: fmt kvstore cli

fmt:
	gofmt -w .

kvstore: gen-protos
	go build -o bin/kvstore ./cmd/kvstore

gen-protos:
	./scripts/gen_protos.sh

cli: gen-protos
	go build -o bin/cli ./cmd/cli

dbuild-server:
	docker build -f Dockerfile-kvstore -t kvstore-server .

dbuild-cli:
	docker build -f Dockerfile-cli -t kvstore-cli .


