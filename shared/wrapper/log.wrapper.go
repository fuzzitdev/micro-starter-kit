package wrapper

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

// LogWrapper is a handler wrapper to log Requests
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Infof("Request: %s", req.Method())
		return fn(ctx, req, rsp)
	}
}

// LogHandlerWrapper is a handler wrapper to log Requests with Context
func LogHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		md, _ := metadata.FromContext(ctx)
		log.WithFields(map[string]interface{}{
			"ctx":    md,
			"method": req.Method(),
		}).Infof("Serving request")

		err := fn(ctx, req, rsp)

		return err
	}
}

type clientLogWrapper struct {
	client.Client
}

func (l *clientLogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// NewClientLogWrapper is a client wrapper to log Requests with metadata
func NewClientLogWrapper(c client.Client) client.Client {
	return &clientLogWrapper{c}
}
