package factory_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/bhagyaraj1208117/andes-abc-1/core/check"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/config"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/libp2p/discovery"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/libp2p/discovery/factory"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/mock"
	"github.com/bhagyaraj1208117/andes-communication-go/testscommon"
	"github.com/stretchr/testify/assert"
)

func TestNewPeerDiscoverer_NilLoggerShouldError(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context:            context.Background(),
		Host:               &mock.ConnectableHostStub{},
		Sharder:            &mock.SharderStub{},
		P2pConfig:          config.P2PConfig{},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             nil,
	}
	pDiscoverer, err := factory.NewPeerDiscoverer(args)
	assert.Equal(t, p2p.ErrNilLogger, err)
	assert.Nil(t, pDiscoverer)
}
func TestNewPeerDiscoverer_NoDiscoveryEnabledShouldRetNullDiscoverer(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context: context.Background(),
		Host:    &mock.ConnectableHostStub{},
		Sharder: &mock.SharderStub{},
		P2pConfig: config.P2PConfig{
			KadDhtPeerDiscovery: config.KadDhtPeerDiscoveryConfig{
				Enabled: false,
			},
		},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             &testscommon.LoggerStub{},
	}
	pDiscoverer, err := factory.NewPeerDiscoverer(args)
	_, ok := pDiscoverer.(*discovery.NilDiscoverer)

	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestNewPeerDiscoverer_ListsSharderShouldWork(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context: context.Background(),
		Host:    &mock.ConnectableHostStub{},
		Sharder: &mock.KadSharderStub{},
		P2pConfig: config.P2PConfig{
			KadDhtPeerDiscovery: config.KadDhtPeerDiscoveryConfig{
				Enabled:                          true,
				RefreshIntervalInSec:             1,
				RoutingTableRefreshIntervalInSec: 300,
				Type:                             "legacy",
			},
			Sharding: config.ShardingConfig{
				Type: p2p.ListsSharder,
			},
		},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             &testscommon.LoggerStub{},
	}

	pDiscoverer, err := factory.NewPeerDiscoverer(args)
	assert.Equal(t, "*discovery.continuousKadDhtDiscoverer", fmt.Sprintf("%T", pDiscoverer))

	assert.NotNil(t, pDiscoverer)
	assert.Nil(t, err)
}

func TestNewPeerDiscoverer_OptimizedKadDhtShouldWork(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context: context.Background(),
		Host:    &mock.ConnectableHostStub{},
		Sharder: &mock.KadSharderStub{},
		P2pConfig: config.P2PConfig{
			KadDhtPeerDiscovery: config.KadDhtPeerDiscoveryConfig{
				Enabled:                          true,
				RefreshIntervalInSec:             1,
				RoutingTableRefreshIntervalInSec: 300,
				Type:                             "optimized",
			},
			Sharding: config.ShardingConfig{
				Type: p2p.ListsSharder,
			},
		},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             &testscommon.LoggerStub{},
	}
	pDiscoverer, err := factory.NewPeerDiscoverer(args)

	assert.Nil(t, err)
	assert.NotNil(t, pDiscoverer)
	assert.Equal(t, "optimized kad-dht discovery", pDiscoverer.Name())
}

func TestNewPeerDiscoverer_UnknownSharderShouldErr(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context: context.Background(),
		Host:    &mock.ConnectableHostStub{},
		Sharder: &mock.SharderStub{},
		P2pConfig: config.P2PConfig{
			KadDhtPeerDiscovery: config.KadDhtPeerDiscoveryConfig{
				Enabled:              true,
				RefreshIntervalInSec: 1,
			},
			Sharding: config.ShardingConfig{
				Type: "unknown",
			},
		},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             &testscommon.LoggerStub{},
	}

	pDiscoverer, err := factory.NewPeerDiscoverer(args)

	assert.True(t, check.IfNil(pDiscoverer))
	assert.True(t, errors.Is(err, p2p.ErrInvalidValue))
}

func TestNewPeerDiscoverer_UnknownKadDhtShouldErr(t *testing.T) {
	t.Parallel()

	args := factory.ArgsPeerDiscoverer{
		Context: context.Background(),
		Host:    &mock.ConnectableHostStub{},
		Sharder: &mock.SharderStub{},
		P2pConfig: config.P2PConfig{
			KadDhtPeerDiscovery: config.KadDhtPeerDiscoveryConfig{
				Enabled:                          true,
				RefreshIntervalInSec:             1,
				RoutingTableRefreshIntervalInSec: 300,
				Type:                             "unknown",
			},
			Sharding: config.ShardingConfig{
				Type: p2p.ListsSharder,
			},
		},
		ConnectionsWatcher: &mock.ConnectionsWatcherStub{},
		Logger:             &testscommon.LoggerStub{},
	}

	pDiscoverer, err := factory.NewPeerDiscoverer(args)

	assert.True(t, errors.Is(err, p2p.ErrInvalidValue))
	assert.True(t, check.IfNil(pDiscoverer))
}
