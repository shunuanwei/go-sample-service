package grpcservers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/api/proto"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *ReviewsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterReviewsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewReviewsServer, CreateInitServersFn)
