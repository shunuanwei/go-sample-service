package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/sdgmf/go-sample-service/api/proto"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/grpc"
)

func NewReviewsClient(client *grpc.Client) (proto.ReviewsClient, error) {
	conn, err := client.Dial("Reviews")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewReviewsClient(conn)

	return c, nil
}
