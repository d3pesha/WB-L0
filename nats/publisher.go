package nats

import (
	"math/rand"
	"strconv"
	"time"
	"wb/models"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomOrder() *models.Order {
	orderUID := randomString(10)
	trackNumber := "WB" + randomString(10)
	timeNow := time.Now().Format(time.RFC3339)

	order := models.Order{
		OrderUID:    orderUID,
		TrackNumber: trackNumber,
		Entry:       randomString(4),
		Delivery: models.Delivery{
			Name:    randomString(randomInt(3, 8)) + " " + randomString(randomInt(5, 12)),
			Phone:   "+" + strconv.Itoa(randomInt(1, 999)) + strconv.Itoa(randomInt(1000000000, 9999999999)),
			Zip:     strconv.Itoa(randomInt(1000000, 9999999)),
			City:    randomString(randomInt(3, 19)),
			Address: randomString(randomInt(5, 9)) + " " + randomString(randomInt(4, 7)) + " " + strconv.Itoa(randomInt(1, 300)),
			Region:  randomString(randomInt(3, 15)),
			Email:   randomString(randomInt(4, 15)) + "@gmail.com",
		},
		Payment: models.Payment{
			Transaction:  orderUID,
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       randomInt(100, 9999),
			PaymentDt:    int(time.Now().Unix()),
			Bank:         randomString(randomInt(4, 13)),
			DeliveryCost: randomInt(100, 9999),
			GoodsTotal:   randomInt(100, 999),
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      randomInt(1000000, 9999999),
				TrackNumber: trackNumber,
				Price:       randomInt(100, 9999),
				Rid:         randomString(17) + "test",
				Name:        randomString(randomInt(4, 13)),
				Sale:        randomInt(0, 70),
				Size:        randomString(1),
				TotalPrice:  randomInt(100, 999),
				NmID:        randomInt(1000000, 9999999),
				Brand:       randomString(randomInt(4, 18)),
				Status:      randomInt(100, 400),
			},
			{
				ChrtID:      randomInt(1000000, 9999999),
				TrackNumber: trackNumber,
				Price:       randomInt(100, 9999),
				Rid:         randomString(17) + "test",
				Name:        randomString(randomInt(4, 13)),
				Sale:        randomInt(0, 70),
				Size:        randomString(1),
				TotalPrice:  randomInt(100, 999),
				NmID:        randomInt(1000000, 9999999),
				Brand:       randomString(randomInt(4, 18)),
				Status:      randomInt(100, 400),
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       timeNow,
		OofShard:          "1",
	}

	return &order
}
