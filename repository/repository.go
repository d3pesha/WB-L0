package repository

import (
	"wb/config"
)

type Database interface {
	Save(order config.Order) error
	Load(uid string) (config.Order, bool)
}

/*type OrderService interface {
	Save(order config.Order) error
}*/
