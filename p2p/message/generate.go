//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/bhagyaraj1208117/protobuf/protobuf  --gogoslick_out=. peerShardMessage.proto

package message
