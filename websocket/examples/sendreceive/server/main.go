package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bhagyaraj1208117/andes-abc-1/marshal/factory"
	"github.com/bhagyaraj1208117/andes-communication-go/websocket/data"
	factoryHost "github.com/bhagyaraj1208117/andes-communication-go/websocket/factory"
	logger "github.com/bhagyaraj1208117/andes-logger-xyz"
)

var (
	marshaller, _ = factory.NewMarshalizer("json")
	log           = logger.GetOrCreate("server")
	url           = "localhost:12345"
)

func main() {
	args := factoryHost.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        url,
			Mode:                       data.ModeServer,
			RetryDurationInSec:         1,
			WithAcknowledge:            true,
			BlockingAckOnError:         false,
			DropMessagesIfNoConnection: false,
			AcknowledgeTimeoutInSec:    10,
		},
		Marshaller: marshaller,
		Log:        log,
	}

	wsServer, err := factoryHost.CreateWebSocketHost(args)
	if err != nil {
		log.Error("cannot create WebSocket server", "error", err)
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	timer := time.NewTimer(100 * time.Millisecond)
	defer timer.Stop()
	count := 0

	func() {
		for {
			message := fmt.Sprintf("message #%d", count)
			err = wsServer.Send([]byte(message), "test")
			if err == nil {
				count++
				log.Info("message sent to clients", "count", count)
			}
			timer.Reset(100 * time.Millisecond)

			select {
			case <-timer.C:
			case <-interrupt:
				return
			}
		}
	}()

	err = wsServer.Close()
	log.LogIfError(err)
}
