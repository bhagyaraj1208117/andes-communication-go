package mock

import "github.com/bhagyaraj1208117/andes-abc-1/core"

// ConnectionsWatcherStub -
type ConnectionsWatcherStub struct {
	NewKnownConnectionCalled func(pid core.PeerID, connection string)
	CloseCalled              func() error
}

// NewKnownConnection -
func (stub *ConnectionsWatcherStub) NewKnownConnection(pid core.PeerID, connection string) {
	if stub.NewKnownConnectionCalled != nil {
		stub.NewKnownConnectionCalled(pid, connection)
	}
}

// Close -
func (stub *ConnectionsWatcherStub) Close() error {
	if stub.CloseCalled != nil {
		return stub.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (stub *ConnectionsWatcherStub) IsInterfaceNil() bool {
	return stub == nil
}
