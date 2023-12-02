package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"wb/models"
)

type ConfigDB struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

type Postgres struct {
	db *sqlx.DB
}

func NewPostgresConfig() *ConfigDB {
	err := godotenv.Load("C:\\Users\\Reflex1on\\GolandProjects\\wb\\.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return &ConfigDB{
		DBHost:     host,
		DBPort:     port,
		DBUsername: username,
		DBPassword: password,
		DBName:     dbname,
	}
}

func Connect(config *ConfigDB) Postgres {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
	)

	db, err := sqlx.Connect("pgx", connStr)

	// defer db.Close() sql: database is closed
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return Postgres{db: db}
}

// сохраняет заказ в базе данных PostgreSQL.
func (p *Postgres) Save(order *models.Order) error {
	config := NewPostgresConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
	)

	tx, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
    INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id,
        delivery_service, shardkey, sm_id, date_created, oof_shard)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID,
		order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
    INSERT INTO delivery (name, phone, zip, city, address, region, email, order_uid)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address,
		order.Delivery.Region, order.Delivery.Email, order.OrderUID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
    INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt,
        bank, delivery_cost, goods_total, custom_fee, order_uid)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		_, err = tx.Exec(`
        INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size,
            total_price, nm_id, brand, status, order_uid)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size,
			item.TotalPrice, item.NmID, item.Brand, item.Status, order.OrderUID)
		if err != nil {
			return err
		}
	}

	return nil
}

// загружает заказ по его уникальному идентификатору (UID).
func (p *Postgres) LoadOrderByUID(uid string) (*models.Order, bool) { //*
	orderQuery := "SELECT * FROM orders WHERE order_uid = $1"
	var order models.Order

	err := p.db.Get(&order, orderQuery, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		log.Printf("Error fetching order: %v", err)
		return nil, false
	}

	deliveryQuery := "SELECT * FROM delivery WHERE order_uid = $1"
	var delivery models.Delivery
	err = p.db.Get(&delivery, deliveryQuery, uid)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error fetching order: %v", err)
		return nil, false
	}

	paymentQuery := "SELECT * FROM payments WHERE order_uid = $1"
	var payment models.Payment
	err = p.db.Get(&payment, paymentQuery, uid)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error fetching order: %v", err)
		return nil, false
	}

	itemsQuery := "SELECT * FROM items WHERE order_uid = $1"
	var items []models.Item
	err = p.db.Select(&items, itemsQuery, uid)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error fetching order: %v", err)
		return nil, false
	}

	order.Delivery = delivery
	order.Payment = payment
	order.Items = items

	return &order, true
}

// загружает все заказы из базы данных PostgreSQL.
func (p *Postgres) LoadAll() (*[]*models.Order, bool) {
	var orders []*models.Order

	err := p.db.Select(&orders, "SELECT * FROM orders;")
	if err != nil {
		log.Fatal(err)
	}
	if len(orders) == 0 {
		return nil, false
	}
	for i, order := range orders {

		loadedOrder, ok := p.LoadOrderByUID(order.OrderUID)
		if ok {
			orders[i] = loadedOrder
		}
	}
	return &orders, true
}
