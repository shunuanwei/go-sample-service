// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-sample-service/internal/app/details"
	"github.com/sdgmf/go-sample-service/internal/app/details/controllers"
	"github.com/sdgmf/go-sample-service/internal/app/details/grpcservers"
	"github.com/sdgmf/go-sample-service/internal/app/details/repositories"
	"github.com/sdgmf/go-sample-service/internal/app/details/services"
	"github.com/sdgmf/go-sample-service/internal/pkg/app"
	"github.com/sdgmf/go-sample-service/internal/pkg/config"
	"github.com/sdgmf/go-sample-service/internal/pkg/consul"
	"github.com/sdgmf/go-sample-service/internal/pkg/database"
	"github.com/sdgmf/go-sample-service/internal/pkg/jaeger"
	"github.com/sdgmf/go-sample-service/internal/pkg/log"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/grpc"
	"github.com/sdgmf/go-sample-service/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	details.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
