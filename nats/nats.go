package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
	"wb/config"
	"wb/database/cache"
	"wb/database/db"
)

var CacheInstance = cache.New()

func Pub() {
	clusterID := "test-cluster"
	clientID := "WBpublisher"

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Error connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	postgresInstance := &db.Postgres{}
	sub, err := sc.Subscribe("WBorder", func(msg *stan.Msg) {
		var order config.Order // **
		err = json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("Error unmarshalling order:", err)
			return
		}

		if err != nil {
			log.Println("Error inserting order into database:", err)
		}

	})

	if err != nil {
		log.Fatalf("Error subscribing to WBorder channel: %v", err)
	}

	defer sub.Close()

	for {
		order := RandomOrder()
		savedInCache := make(chan struct{}, 1)
		err = postgresInstance.Save(order)
		go func(order config.Order) {
			err = CacheInstance.Save(order)
			if err != nil {
				log.Println("Error inserting order into cache: ", err)
			}

			savedInCache <- struct{}{}
		}(*order)

		<-savedInCache

		msg, err := json.Marshal(order)
		if err != nil {
			log.Fatalf("Error with marshal order to JSON: %v", err)
		}

		err = sc.Publish("WBorder", msg)
		if err != nil {
			log.Fatalf("Can't publish message into NATS: %v, ", err)
		}

		time.Sleep(10 * time.Second)
	}

}
