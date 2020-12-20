SRC_DIR="proto/"
DST_DIR="protogen/"

protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR proto/KVStore.proto