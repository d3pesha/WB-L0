package database

import (
	"fmt"
	"log"
	"wb/database/cache"
	"wb/database/db"
	"wb/models"
)

type Memory struct {
	postgres db.Postgres
	Cache    cache.Cache
}

func New() Memory {
	m := Memory{
		postgres: db.Connect(db.NewPostgresConfig()),
		Cache:    cache.New(),
	}
	m.DBtoMem()
	return m
}

func (m *Memory) DBtoMem() {
	orders, ok := m.postgres.LoadAll()
	if !ok {
		return
	}
	for _, orderPtr := range *orders {
		if orderPtr == nil || orderPtr.OrderUID == "" {
			log.Println("Found nil order while loading from database.")
			continue
		}
		order := *orderPtr
		if err := m.Cache.Save(order); err != nil {
			log.Println("Failed to save order to cache:", err)
		}
	}
	return
}

func (m *Memory) Save(order models.Order) error {
	fmt.Println("saved id : " + order.OrderUID)
	if err := m.Cache.Save(order); err != nil {
		return err
	}
	if err := m.postgres.Save(&order); err != nil {
		return err
	}
	return nil
}

func (m *Memory) Load(uid string) (models.Order, bool) {
	return m.Cache.Load(uid)
}
