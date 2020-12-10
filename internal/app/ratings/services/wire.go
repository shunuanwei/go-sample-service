// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/ratings/repositories"
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

func CreateRatingsService(cf string, sto repositories.RatingsRepository) (RatingsService, error) {
	panic(wire.Build(testProviderSet))
}
