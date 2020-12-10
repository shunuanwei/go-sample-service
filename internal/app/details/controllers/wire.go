// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/details/repositories"
	"github.com/sdgmf/go-sample-service/internal/app/details/services"
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

func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(testProviderSet))
}
