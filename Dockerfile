FROM golang:alpine

RUN apk add --no-cache bash make protoc git \
    && go get github.com/golang/protobuf/protoc-gen-go \
    && cp /go/bin/protoc-gen-go /usr/bin/

WORKDIR /build

COPY . .

RUN make kvstore

CMD ["/build/bin/kvstore"]
