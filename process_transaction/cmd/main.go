package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/Daniel-Vinicus/imersaofc5/adapter/broker/kafka"
	"github.com/Daniel-Vinicus/imersaofc5/adapter/factory"
	"github.com/Daniel-Vinicus/imersaofc5/adapter/presenter/transaction"
	"github.com/Daniel-Vinicus/imersaofc5/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	transactionRepository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	// Other thread, in parallel reading the messages
	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(transactionRepository, producer, "transactions_result")

	// Processing the messages already read
	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
