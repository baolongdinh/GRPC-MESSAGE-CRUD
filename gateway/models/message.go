package models

import (
	"time"

	messagegrpc "gateway-service/proto/message"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;"`
}

type SendMessageResquest struct {
	Message string `json:"message"`
}

type SendMessageResponse struct {
	Message string `json:"message"`
}

type GetMessageListRequest struct {
}

type GetMessageListResponse struct {
	Messages []*Message
}

type GetMessageByIdRequest struct {
	Id string `json:"id" validate:"required" default:"00000000-0000-0000-0000-000000000000"`
}

type GetMessageByIdResponse struct {
	Message string `json:"message"`
}

type UpdateMessageByIdRequest struct {
	Id      string `json:"id" validate:"required" default:"00000000-0000-0000-0000-000000000000"`
	Message string `json:"message"`
}

type UpdateMessageByIdResponse struct {
	Message string `json:"message"`
}

type DeleteMessageByIdRequest struct {
	Id string `json:"id" validate:"required" default:"00000000-0000-0000-0000-000000000000"`
}

type DeleteMessageByIdResponse struct {
	Respone string `json:"respone"`
}

func ConvertToMessage(m *messagegrpc.Message) *Message {
	return &Message{
		Id:        uuid.MustParse(m.Id),
		Message:   m.Message,
		CreatedAt: m.CreatedAt.AsTime(),
	}
}

func ConvertToListMessage(messages []*messagegrpc.Message) []*Message {

	var result []*Message

	for _, m := range messages {
		result = append(result, ConvertToMessage(m))
	}
	return result
}

func ConvertToRpcMessage(m *Message) *messagegrpc.Message {
	createdAt := timestamppb.New(m.CreatedAt)

	return &messagegrpc.Message{
		Id:        m.Id.String(),
		Message:   m.Message,
		CreatedAt: createdAt,
	}
}

func ConvertToRpcListMessage(listMessage []*Message) []*messagegrpc.Message {
	var result []*messagegrpc.Message
	for _, m := range listMessage {
		result = append(result, ConvertToRpcMessage(m))
	}
	return result
}
