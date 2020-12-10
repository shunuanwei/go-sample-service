package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/sdgmf/go-sample-service/api/proto"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/grpc"
)

func NewDetailsClient(client *grpc.Client) (proto.DetailsClient, error) {
	conn, err := client.Dial("Details")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewDetailsClient(conn)

	return c, nil
}
