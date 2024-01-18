package peerDiscovery

import (
	"bytes"
	"sync"

	"github.com/bhagyaraj1208117/andes-abc-1/core"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p"
)

// MessageProcesssor -
type MessageProcesssor struct {
	RequiredValue   []byte
	chanDone        chan struct{}
	mutDataReceived sync.Mutex
	wasDataReceived bool
}

// NewMessageProcessor -
func NewMessageProcessor(chanDone chan struct{}, requiredVal []byte) *MessageProcesssor {
	return &MessageProcesssor{
		RequiredValue: requiredVal,
		chanDone:      chanDone,
	}
}

// ProcessReceivedMessage -
func (mp *MessageProcesssor) ProcessReceivedMessage(message p2p.MessageP2P, _ core.PeerID, _ p2p.MessageHandler) error {
	if bytes.Equal(mp.RequiredValue, message.Data()) {
		mp.mutDataReceived.Lock()
		mp.wasDataReceived = true
		mp.mutDataReceived.Unlock()

		mp.chanDone <- struct{}{}
	}

	return nil
}

// WasDataReceived -
func (mp *MessageProcesssor) WasDataReceived() bool {
	mp.mutDataReceived.Lock()
	defer mp.mutDataReceived.Unlock()

	return mp.wasDataReceived
}

// IsInterfaceNil returns true if there is no value under the interface
func (mp *MessageProcesssor) IsInterfaceNil() bool {
	return mp == nil
}