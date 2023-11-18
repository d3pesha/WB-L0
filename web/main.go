package main

import (
	"log"
	"wb/cmd/app"
	"wb/nats"
)

func main() {
	go nats.Pub()
	// Создаем экземпляр приложения и запускаем его
	test := app.New()
	if err := test.Run(); err != nil {
		log.Printf("Error running the server: %v", err)
	}

}
