package main

import (
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/jackc/pgx/v5/stdlib"

	"balance-service/internal/database"
	"balance-service/internal/event/handler"
	"balance-service/internal/usecase/get_balance"
	"balance-service/internal/web"
	"balance-service/internal/web/webserver"
	"balance-service/pkg/kafka"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@balance-db:5432/balance?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDB := database.NewBalanceDB(db)
	balanceUpdatedHandler := handler.NewBalanceUpdatedHandler(balanceDB)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "balance-service",
	}
	consumer := kafka.NewConsumer(&configMap, []string{"balances"})

	msgChan := make(chan *ckafka.Message)
	go consumer.Consume(msgChan)
	go func() {
		for msg := range msgChan {
			if err := balanceUpdatedHandler.Handle(msg.Value); err != nil {
				fmt.Println("error handling message:", err)
			}
		}
	}()

	getBalanceUseCase := get_balance.NewGetBalanceUseCase(balanceDB)
	balanceHandler := web.NewWebBalanceHandler(*getBalanceUseCase)

	ws := webserver.NewWebServer(":3003")
	ws.AddHandler("/balances/{accountId}", balanceHandler.GetBalance)

	fmt.Println("balance-service running")
	ws.Start()
}
