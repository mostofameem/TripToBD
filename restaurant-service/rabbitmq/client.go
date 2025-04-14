package rabbitmq

import (
	"restaurant-service/config"
	"time"

	rmq "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	connectionManager *ConnectionManager
	consumerManager   *ConsumerManager
	publisher         *Publisher
}

var client *Client

func GetClient() *Client {
	return client
}

func NewClient() *Client {
	conf := config.GetConfig()
	c := &Client{
		connectionManager: NewConnectionManager(
			conf.RabbitmqURL,
			time.Duration(conf.RmqReconnectDelay)*time.Second,
		),
		consumerManager: NewConsumerManager(),
		publisher:       NewPublisher(),
	}

	// consumer manager listener for connection events
	c.connectionManager.OnConnect(c.consumerManager.HandleConnect)
	c.connectionManager.OnDisconnect(c.consumerManager.HandleDisconnect)

	// publisher listener for connection events
	c.connectionManager.OnConnect(c.publisher.HandleConnect)
	c.connectionManager.OnDisconnect(c.publisher.HandleDisconnect)

	return c
}

func (c *Client) Start() {
	c.connectionManager.Start()
	c.publisher.Start()
}

func (c *Client) Stop() {
	c.connectionManager.Stop()
	c.publisher.Stop()
}

func (c *Client) AddConsumer(consumer ConsumerOption) {
	c.consumerManager.Add(consumer)
}

func (c *Client) Publish(msg Message) {
	c.publisher.Publish(msg)
}

func (c *Client) GetChannel() *rmq.Channel {
	return c.connectionManager.getChannel()
}