package mock

import "github.com/bhagyaraj1208117/andes-abc-1/core"

// PeersRatingHandlerStub -
type PeersRatingHandlerStub struct {
	IncreaseRatingCalled           func(pid core.PeerID)
	DecreaseRatingCalled           func(pid core.PeerID)
	GetTopRatedPeersFromListCalled func(peers []core.PeerID, numOfPeers int) []core.PeerID
}

// IncreaseRating -
func (stub *PeersRatingHandlerStub) IncreaseRating(pid core.PeerID) {
	if stub.IncreaseRatingCalled != nil {
		stub.IncreaseRatingCalled(pid)
	}
}

// DecreaseRating -
func (stub *PeersRatingHandlerStub) DecreaseRating(pid core.PeerID) {
	if stub.DecreaseRatingCalled != nil {
		stub.DecreaseRatingCalled(pid)
	}
}

// GetTopRatedPeersFromList -
func (stub *PeersRatingHandlerStub) GetTopRatedPeersFromList(peers []core.PeerID, numOfPeers int) []core.PeerID {
	if stub.GetTopRatedPeersFromListCalled != nil {
		return stub.GetTopRatedPeersFromListCalled(peers, numOfPeers)
	}

	return peers
}

// IsInterfaceNil returns true if there is no value under the interface
func (stub *PeersRatingHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}