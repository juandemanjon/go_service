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
go mod init github/juandemanjon/go_service
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## protoc
```
protoc --go_out=paths=source_relative:./gen -I. chat.proto
protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative chat.proto
```