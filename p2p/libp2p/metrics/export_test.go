package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/bhagyaraj1208117/andes-abc-1/core"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p"
	"github.com/bhagyaraj1208117/andes-storage-go/timecache"
)

const MinTimeToLive = minTimeToLive

// NewPrintConnectionsWatcherWithHandler -
func NewPrintConnectionsWatcherWithHandler(timeToLive time.Duration, handler func(pid core.PeerID, connection string, log p2p.Logger)) (*printConnectionsWatcher, error) {
	if timeToLive < minTimeToLive {
		return nil, fmt.Errorf("%w in NewPrintConnectionsWatcher, got: %d, minimum: %d", ErrInvalidValueForTimeToLiveParam, timeToLive, minTimeToLive)
	}

	pcw := &printConnectionsWatcher{
		timeToLive:   timeToLive,
		timeCacher:   timecache.NewTimeCache(timeToLive),
		printHandler: handler,
	}

	ctx, cancel := context.WithCancel(context.Background())
	pcw.cancel = cancel
	go pcw.doSweep(ctx)

	return pcw, nil
}

func LogPrintHandler(pid core.PeerID, connection string, log p2p.Logger) {
	logPrintHandler(pid, connection, log)
}

func (pcw *printConnectionsWatcher) GoRoutineClosed() bool {
	return pcw.goRoutineClosed.IsSet()
}
