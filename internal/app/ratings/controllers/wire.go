// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/ratings/repositories"
	"github.com/sdgmf/go-sample-service/internal/app/ratings/services"
	"github.com/sdgmf/go-sample-service/internal/pkg/config"
	"github.com/sdgmf/go-sample-service/internal/pkg/database"
	"github.com/sdgmf/go-sample-service/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateRatingsController(cf string, sto repositories.RatingsRepository) (*RatingsController, error) {
	panic(wire.Build(testProviderSet))
}
