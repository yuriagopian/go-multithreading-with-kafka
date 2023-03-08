package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/devfullcycle/gointesivo2/internal/infra/database"
	"github.com/devfullcycle/gointesivo2/internal/usecase"
	"github.com/devfullcycle/gointesivo2/pkg/kafka"
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
	go kafka.Consume(topics, servers, msgChanKafka)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	for msg := range msgChan {
		var OrderInputDto usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDto)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(OrderInputDto)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Kafka has processed order %s\n", outputDto.ID)
	}
}
