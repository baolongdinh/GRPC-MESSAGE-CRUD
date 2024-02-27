package main

import (
	client "gateway-service/client"
	grpc "gateway-service/client/grpc"
	config "gateway-service/config"
	handlers "gateway-service/handlers"

	middleware "gateway-service/middlewares"
	routes "gateway-service/routes"

	"github.com/labstack/echo/v4"
)

var (
	server      *echo.Echo
	proxyLogger middleware.ProxyLogger

	messageClient  client.MessageClient
	messageHandler handlers.MessageHandler
	messageRoute   routes.MessageRoute
)

func init() {
	appConfig := config.MustLoadConfig(".", "app")

	messageClient = grpc.MustNewDatasetGrpcClient(appConfig.MessageServiceAddress)
	messageHandler = handlers.NewMessageHandler(messageClient)
	messageRoute = routes.NewMessageRoute(messageHandler)

}
