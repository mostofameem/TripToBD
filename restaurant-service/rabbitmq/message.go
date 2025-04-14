package rabbitmq

import (
	"restaurant-service/logger"
)

type Message struct {
	Exchange   string
	RoutingKey string
	Message    []byte
}

func NewMessage(exchange, routingKey string, data interface{}) Message {
	dataStr := logger.ConvertToJson(data)
	return Message{
		Exchange:   exchange,
		RoutingKey: routingKey,
		Message:    []byte(dataStr),
	}
}
