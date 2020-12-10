package grpcservers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/api/proto"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *RatingsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterRatingsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewRatingsServer, CreateInitServersFn)
