package grpc

import (
	"fmt"
	client "gateway-service/client"
	models "gateway-service/models"
	messagegrpc "gateway-service/proto/message"
	utils "gateway-service/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MessageGrpcClient struct {
	remoteAddress     string
	MessageGrpcClient messagegrpc.MessageServiceClient
}

func MustNewDatasetGrpcClient(remoteAddress string) client.MessageClient {
	datasetServiceConfig := fmt.Sprintf(`{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": "%s"
	}
	}`,
		messagegrpc.MessageService_ServiceDesc.ServiceName)

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(datasetServiceConfig),
	}

	conn, err := grpc.Dial(remoteAddress, options...)
	if err != nil {
		log.Fatal(err)
	}

	client := messagegrpc.NewMessageServiceClient(conn)
	return &MessageGrpcClient{
		remoteAddress:     remoteAddress,
		MessageGrpcClient: client,
	}
}

func (c *MessageGrpcClient) SendMessage(ctx echo.Context, req models.SendMessageResquest) error {
	res, err := c.MessageGrpcClient.SendMessage(ctx.Request().Context(), &messagegrpc.SendMessageResquest{
		Message: req.Message,
	})

	if err != nil {
		return utils.HandleGrpcStatus(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "", res.Message)
}

func (c *MessageGrpcClient) GetMessageList(ctx echo.Context, req models.GetMessageListRequest) error {
	res, err := c.MessageGrpcClient.GetMessageList(ctx.Request().Context(), &messagegrpc.GetMessageListRequest{})

	if err != nil {
		return utils.HandleGrpcStatus(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "get messages list sucessfully", models.ConvertToListMessage(res.Messages))
}

func (c *MessageGrpcClient) GetMessageById(ctx echo.Context, req models.GetMessageByIdRequest) error {
	res, err := c.MessageGrpcClient.GetMessageById(ctx.Request().Context(), &messagegrpc.GetMessageByIdRequest{
		Id: req.Id,
	})

	if err != nil {
		return utils.HandleGrpcStatus(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "get message sucessfully", models.ConvertToMessage(res.Message))
}

func (c *MessageGrpcClient) UpdateMessageById(ctx echo.Context, req models.UpdateMessageByIdRequest) error {
	res, err := c.MessageGrpcClient.UpdateMessageById(ctx.Request().Context(), &messagegrpc.UpdateMessageByIdRequest{
		Id:      req.Id,
		Message: req.Message,
	})

	if err != nil {
		return utils.HandleGrpcStatus(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "updated message sucessfully", models.ConvertToMessage(res.Message))
}

func (c *MessageGrpcClient) DeleteMessageById(ctx echo.Context, req models.DeleteMessageByIdRequest) error {
	_, err := c.MessageGrpcClient.DeleteMessageById(ctx.Request().Context(), &messagegrpc.DeleteMessageByIdRequest{
		Id: req.Id,
	})

	if err != nil {
		return utils.HandleGrpcStatus(ctx, err)
	}

	return utils.Response(ctx, http.StatusOK, "delete message sucessfully", nil)
}
