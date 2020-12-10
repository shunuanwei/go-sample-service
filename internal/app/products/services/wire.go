// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/api/proto"
	"github.com/sdgmf/go-sample-service/internal/pkg/config"
	"github.com/sdgmf/go-sample-service/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ProviderSet,
)

func CreateProductsService(cf string,
	detailsSvc proto.DetailsClient,
	ratingsSvc proto.RatingsClient,
	reviewsSvc proto.ReviewsClient) (ProductsService, error) {
	panic(wire.Build(testProviderSet))
}
