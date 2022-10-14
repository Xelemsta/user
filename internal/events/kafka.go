package events

import (
	"context"
	"encoding/json"
	"os"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	CreateUserAction Action = "createUser"
	UpdateUser       Action = "updateUser"
	DeleteUser       Action = "deleteAction"
)

// Action refers to user action
// we want to notify
type Action string

// Data that will be send
// inside the topic
type Data struct {
	UserID *string
	Action Action
}

// Produce produces "KAFKA_TOPIC" topic
// environment variable must contains
// "KAFKA_TOPIC" and "BROKER_KAFKA" values
// in order to operate properly
func Produce(ctx context.Context, data *Data) {
	if os.Getenv("KAFKA_TOPIC") == "" {
		log.Error(`no KAFKA_TOPIC found to produce`)
		return
	}

	if os.Getenv("BROKER_KAFKA") == "" {
		log.Error(`no BROKER_KAFKA found to produce`)
		return
	}

	if data.UserID == nil || *data.UserID == "" {
		log.Error(`user id is mandatory in order to produce topic`)
	}

	log.Infof(`producing topic %s`, os.Getenv("KAFKA_TOPIC"))

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("BROKER_KAFKA")},
		Topic:   os.Getenv("KAFKA_TOPIC"),
	})

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Errorf(`error while marshalling data: %+v`, err)
		return
	}

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte("data"),
		Value: bytes,
	})

	if err != nil {
		log.Errorf(`error while sending topic: %+v`, err)
	} else {
		log.Infof(`successfully produced topic %s`, os.Getenv("KAFKA_TOPIC"))
	}
}
