package networksharding_test

import (
	"testing"

	"github.com/bhagyaraj1208117/andes-abc-1/core/check"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/libp2p/networksharding"
	"github.com/stretchr/testify/assert"
)

func TestNilListSharderSharder(t *testing.T) {
	nls := networksharding.NewNilListSharder()

	assert.False(t, check.IfNil(nls))
	assert.Equal(t, 0, len(nls.ComputeEvictionList(nil)))
	assert.False(t, nls.Has("", nil))
	assert.Nil(t, nls.SetPeerShardResolver(nil))
}
