// +build wireinject

package repositories

import (
	"github.com/google/wire"
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

func CreateDetailRepository(f string) (DetailsRepository, error) {
	panic(wire.Build(testProviderSet))
}
