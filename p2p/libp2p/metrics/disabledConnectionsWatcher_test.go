package metrics_test

import (
	"fmt"
	"testing"

	"github.com/bhagyaraj1208117/andes-abc-1/core/check"
	"github.com/bhagyaraj1208117/andes-communication-go/p2p/libp2p/metrics"
	"github.com/stretchr/testify/assert"
)

func TestDisabledConnectionsWatcher_MethodsShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("should have not panic: %v", r))
		}
	}()

	dcw := metrics.NewDisabledConnectionsWatcher()
	assert.False(t, check.IfNil(dcw))
	dcw.NewKnownConnection("", "")
	err := dcw.Close()
	assert.Nil(t, err)
}
