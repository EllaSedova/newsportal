package rpc

import (
	zm "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"

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
		//zm.WithDevel(isDevel),
		zm.WithHeaders(),
		zm.WithSentry(zm.DefaultServerName),
		zm.WithNoCancelContext(),
		//zm.WithMetrics(zm.DefaultServerName),
		//zm.WithTiming(isDevel, allowDebugFn()),
		//zm.WithSQLLogger(dbo.DB, isDevel, allowDebugFn(), allowDebugFn()),
	)

	// services
	rpc.RegisterAll(map[string]zenrpc.Invoker{
		NsRPC: NewRPCService(m),
	})

	return rpc
}
