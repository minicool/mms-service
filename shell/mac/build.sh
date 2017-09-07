#!/usr/bin/env bash

PROTO_PATH="../../proto/"
PROTO_NORMAL="DeviceServer.proto"

GO_BUILD_PATH="../../src"
BIN_PATH="../../bin/"

echo "-------rm bin--------"
rm ${BIN_PATH}/mms-device-service
echo "-------build proto-------"
protoc --proto_path=${PROTO_PATH} --go_out=plugins=grpc:${PROTO_PATH} ${PROTO_PATH}/${PROTO_NORMAL}
echo "-------build go-------"
go build -o ${BIN_PATH}/mms-service ${GO_BUILD_PATH}/main.go
echo "-------build go test-------"
go build -o ${BIN_PATH}/test-mms-client ${GO_BUILD_PATH}/../test/client/main.go
