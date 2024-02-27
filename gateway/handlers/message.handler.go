package handlers

import (
	"gateway-service/client"
	models "gateway-service/models"
	utils "gateway-service/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	client client.MessageClient
}

func NewMessageHandler(messageClient client.MessageClient) MessageHandler {
	return MessageHandler{
		client: messageClient,
	}
}

func (h MessageHandler) SendMessage(ctx echo.Context) error {

	var req models.SendMessageResquest

	if err := ctx.Bind(&req); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	return h.client.SendMessage(ctx, req)
}

func (h MessageHandler) GetMessageList(ctx echo.Context) error {

	var req models.GetMessageListRequest

	if err := ctx.Bind(&req); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	return h.client.GetMessageList(ctx, req)
}

func (h MessageHandler) GetMessageById(ctx echo.Context) error {

	var req models.GetMessageByIdRequest

	Id := ctx.Param("id")

	if _, err := uuid.Parse(Id); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	req.Id = Id
	return h.client.GetMessageById(ctx, req)
}

func (h MessageHandler) UpdateMessageById(ctx echo.Context) error {

	var req models.UpdateMessageByIdRequest

	Id := ctx.Param("id")

	if err := ctx.Bind(&req); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	if _, err := uuid.Parse(Id); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	req.Id = Id

	return h.client.UpdateMessageById(ctx, req)
}

func (h MessageHandler) DeleteMessageById(ctx echo.Context) error {

	var req models.DeleteMessageByIdRequest

	Id := ctx.Param("id")

	if _, err := uuid.Parse(Id); err != nil {
		return utils.Response(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	req.Id = Id

	return h.client.DeleteMessageById(ctx, req)
}
