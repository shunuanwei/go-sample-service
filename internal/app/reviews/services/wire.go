// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/reviews/repositories"
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

func CreateReviewsService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
	panic(wire.Build(testProviderSet))
}
