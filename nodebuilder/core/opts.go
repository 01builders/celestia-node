package core

import (
	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/libs/fxutil"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// WithConnection sets a custom client for core process
func WithConnection(conn *grpc.ClientConn) fx.Option {
	return fxutil.ReplaceAs(conn, new(grpc.ClientConn))
}

//func WithConnection(conn *grpc.ClientConn) fx.Option {
//	return fx.Replace(fx.Annotated{
//		Target: conn,
//		Name:   "coreGRPC",
//	})
//}

// WithHeaderConstructFn sets custom func that creates extended header
func WithHeaderConstructFn(construct header.ConstructFn) fx.Option {
	return fx.Replace(construct)
}
