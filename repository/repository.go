package repository

import (
	"wb/models"
)

type Database interface {
	Save(order models.Order) error
	Load(uid string) (models.Order, bool)
}

/*type OrderService interface {
	Save(order models.Order) error
}*/
