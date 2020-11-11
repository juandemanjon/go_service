# go_service

## install protoc
```
PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
unzip -o $PROTOC_ZIP -d local bin/protoc
unzip -o $PROTOC_ZIP -d local 'include/*'
rm -f $PROTOC_ZIP
```

## install grpcurl
```
GRPCURL_TAR=grpcurl_1.7.0_linux_x86_64.tar.gz
curl -OL https://github.com/fullstorydev/grpcurl/releases/download/v1.7.0/$GRPCURL_TAR
tar -C local/bin -zxvf $GRPCURL_TAR
rm -f $GRPCURL_TAR

```

## path
export GOPATH="/home/runner/go"
export PATH=$PATH:$PWD/local/bin:$GOPATH/bin

## go mod
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

go mod init github.com/juandemanjon/go_service
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

## proto generation
```
protoc --go_out=paths=source_relative:./gen -I proto chat.proto
protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative -I proto chat_service.proto

protoc -I proto --grpc-gateway_out ./gen --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis chat_service.proto

```