// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	detailsOptions, err := details.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	detailsRepository := repositories.NewMysqlDetailsRepository(logger, db)
	detailsService := services.NewDetailService(logger, detailsRepository)
	detailsController := controllers.NewDetailsController(logger, detailsService)
	initControllers := controllers.CreateInitControllersFn(detailsController)
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, client)
	if err != nil {
		return nil, err
	}
	serverOptions, err := grpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	detailsServer, err := grpcservers.NewDetailsServer(logger, detailsService)
	if err != nil {
		return nil, err
	}
	initServers := grpcservers.CreateInitServersFn(detailsServer)
	grpcServer, err := grpc.NewServer(serverOptions, logger, initServers, client, tracer)
	if err != nil {
		return nil, err
	}
	application, err := details.NewApp(detailsOptions, logger, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, repositories.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, details.ProviderSet, controllers.ProviderSet, grpcservers.ProviderSet)
