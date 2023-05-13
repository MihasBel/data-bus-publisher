package rep

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ConsumerCreator interface {
	NewConsumer(*kafka.ConfigMap) (*kafka.Consumer, error)
}
