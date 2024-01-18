package websocket

import (
	"github.com/bhagyaraj1208117/andes-abc-1/core/check"
	"github.com/bhagyaraj1208117/andes-abc-1/marshal"
	"github.com/bhagyaraj1208117/andes-communication-go/websocket/data"
)

type webSocketsPayloadConverter struct {
	marshaller marshal.Marshalizer
}

// NewWebSocketPayloadConverter returns a new instance of websocketPayloadParser
func NewWebSocketPayloadConverter(marshaller marshal.Marshalizer) (*webSocketsPayloadConverter, error) {
	if check.IfNil(marshaller) {
		return nil, data.ErrNilMarshaller
	}

	return &webSocketsPayloadConverter{
		marshaller: marshaller,
	}, nil
}

// ExtractWsMessage will extract the provided payload in a *data.WsMessage
func (wpc *webSocketsPayloadConverter) ExtractWsMessage(payload []byte) (*data.WsMessage, error) {
	wsMessage := &data.WsMessage{}
	err := wpc.marshaller.Unmarshal(wsMessage, payload)
	if err != nil {
		return nil, err
	}

	return wsMessage, nil
}

// ConstructPayload will marshal the provided *data.WsMessage
func (wpc *webSocketsPayloadConverter) ConstructPayload(wsMessage *data.WsMessage) ([]byte, error) {
	return wpc.marshaller.Marshal(wsMessage)
}

// IsInterfaceNil -
func (wpc *webSocketsPayloadConverter) IsInterfaceNil() bool {
	return wpc == nil
}
