package rpc

import (
	zm "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"os"

	"newsportal/pkg/newsportal"
)

//go:generate zenrpc

const (
	NsRPC = "rpc"
)

// New returns new zenrpc Server.
func New(m *newsportal.Manager) zenrpc.Server {
	rpc := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD: true,
		AllowCORS: true,
	})

	rpc.Use(
		zm.WithHeaders(),
		zm.WithSentry(zm.DefaultServerName),
		zm.WithNoCancelContext(),
		zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	)

	// services
	rpc.RegisterAll(map[string]zenrpc.Invoker{
		NsRPC: NewNewsService(m),
	})

	return rpc
}
