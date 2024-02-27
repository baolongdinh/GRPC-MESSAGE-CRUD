package routes

import (
	handlers "gateway-service/handlers"

	"github.com/labstack/echo/v4"
)

type MessageRoute struct {
	messageHandler handlers.MessageHandler
}

func NewMessageRoute(messageHandler handlers.MessageHandler) MessageRoute {
	return MessageRoute{
		messageHandler: messageHandler,
	}
}

func (m MessageRoute) Routes(r *echo.Group) {
	messageGroup := r.Group("/message")
	{

		messageGroup.POST("/send", m.messageHandler.SendMessage)
		messageGroup.GET("", m.messageHandler.GetMessageList)
		messageGroup.GET("/:id", m.messageHandler.GetMessageById)
		messageGroup.PUT("/:id", m.messageHandler.UpdateMessageById)
		messageGroup.DELETE("/:id", m.messageHandler.DeleteMessageById)
	}
}
