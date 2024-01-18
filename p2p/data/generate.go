//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/bhagyaraj1208117/protobuf/protobuf  --gogoslick_out=. topicMessage.proto
package data
