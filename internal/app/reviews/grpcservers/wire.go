// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/reviews/services"
	"github.com/sdgmf/go-sample-service/internal/pkg/config"
	"github.com/sdgmf/go-sample-service/internal/pkg/database"
	"github.com/sdgmf/go-sample-service/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	panic(wire.Build(testProviderSet))
}
