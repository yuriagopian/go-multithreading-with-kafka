package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/devfullcycle/gointesivo2/internal/infra/database"
	"github.com/devfullcycle/gointesivo2/internal/usecase"
	"github.com/devfullcycle/gointesivo2/pkg/kafka"
	"github.com/devfullcycle/gointesivo2/pkg/rabbitmq"
	amp "github.com/rabbitmq/amqp091-go"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() // Executa tudo e depois executa o close

	repository := database.NewOrderRepository(db)

	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"

	fmt.Println("Kafka consumer has started")

	// Consumindo os dados do kafka em uma outra thread
	go kafka.Consume(topics, servers, msgChanKafka)

	// Executa o worker do kafka
	go kafkaWorker(msgChanKafka, usecase)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, usecase)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChan {
		var OrderInputDto usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDto)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(OrderInputDto)

		if err != nil {
			fmt.Println("An Erro has occurred", err)
			// Panic when error stop the connection
			// panic(err)
		} else {
			fmt.Printf("Kafka has processed order %s\n", outputDto.ID)
		}

	}
}

func rabbitmqWorker(msgChan chan amp.Delivery, uc usecase.CalculateFinalPrice) {
	fmt.Println("Rabbitmq worker has started")
	for msg := range msgChan {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &OrderInputDTO)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Printf("Rabbitmq has processed order %s\n", outputDto.ID)
	}
}
