package client

import (
	models "gateway-service/models"

	"github.com/labstack/echo/v4"
)

type MessageClient interface {
	SendMessage(ctx echo.Context, req models.SendMessageResquest) error
	GetMessageList(ctx echo.Context, req models.GetMessageListRequest) error
	GetMessageById(ctx echo.Context, req models.GetMessageByIdRequest) error
	UpdateMessageById(ctx echo.Context, req models.UpdateMessageByIdRequest) error
	DeleteMessageById(ctx echo.Context, req models.DeleteMessageByIdRequest) error
}
