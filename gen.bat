@echo off
echo Generating Go code from proto files...

if not exist protos\gen mkdir protos\gen

protoc -I=protos --go_out=protos/gen --go_opt=paths=source_relative --go-grpc_out=protos/gen --go-grpc_opt=paths=source_relative protos\base.proto protos\client.proto

echo Done!