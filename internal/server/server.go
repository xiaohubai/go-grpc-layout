package server

import (
	"github.com/google/wire"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, consul.NewRegistry)
