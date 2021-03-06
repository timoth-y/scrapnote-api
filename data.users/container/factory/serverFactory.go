package factory

import (
	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"github.com/timoth-y/scrapnote-api/lib.common/server"

	"github.com/timoth-y/scrapnote-api/data.users/api/rpc"
	"github.com/timoth-y/scrapnote-api/data.users/config"
)

func ProvideServer(config config.ServiceConfig, asyncHandler core.Handler, rpcHandler *rpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupGRPC(rpc.ProvideRemoteSetup(rpcHandler))
	srv.SetupAMQP(asyncHandler)
	return srv
}
