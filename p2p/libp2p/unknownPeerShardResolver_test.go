package libp2p_test

import (
	"testing"

	"github.com/bhagyaraj1208117/andes-communication-go/p2p/libp2p"
	"github.com/bhagyaraj1208117/andes-core-go/core"
	"github.com/bhagyaraj1208117/andes-core-go/core/check"
	"github.com/stretchr/testify/assert"
)

func TestUnknownPeerShardResolver_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	upsr := libp2p.NewUnknownPeerShardResolver()
	assert.False(t, check.IfNil(upsr))
}

func TestUnknownPeerShardResolver_GetPeerInfoShouldReturnUnknownId(t *testing.T) {
	t.Parallel()

	upsr := libp2p.NewUnknownPeerShardResolver()
	expectedPeerInfo := core.P2PPeerInfo{
		PeerType: core.UnknownPeer,
		ShardID:  0,
	}

	assert.Equal(t, expectedPeerInfo, upsr.GetPeerInfo(""))
}
