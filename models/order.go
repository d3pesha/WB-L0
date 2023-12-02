package models

type Order struct {
	OrderUID          string   `db:"order_uid" json:"order_uid,omitempty"`
	TrackNumber       string   `db:"track_number" json:"track_number,omitempty"`
	Entry             string   `db:"entry" json:"entry,omitempty"`
	Delivery          Delivery `db:"delivery" json:"delivery"`
	Payment           Payment  `db:"payments" json:"payment"`
	Items             []Item   `db:"items" json:"items,omitempty"`
	Locale            string   `db:"locale" json:"locale,omitempty"`
	InternalSignature string   `db:"internal_signature" json:"internal_signature,omitempty"`
	CustomerID        string   `db:"customer_id" json:"customer_id,omitempty"`
	DeliveryService   string   `db:"delivery_service" json:"delivery_service,omitempty"`
	Shardkey          string   `db:"shardkey" json:"shardkey,omitempty"`
	SmID              int      `db:"sm_id" json:"sm_id,omitempty"`
	DateCreated       string   `db:"date_created" json:"date_created,omitempty"`
	OofShard          string   `db:"oof_shard" json:"oof_shard,omitempty"`
}

type Delivery struct {
	Name     string `db:"name" json:"name,omitempty"`
	Phone    string `db:"phone" json:"phone,omitempty"`
	Zip      string `db:"zip" json:"zip,omitempty"`
	City     string `db:"city" json:"city,omitempty"`
	Address  string `db:"address" json:"address,omitempty"`
	Region   string `db:"region" json:"region,omitempty"`
	Email    string `db:"email" json:"email,omitempty"`
	OrderUID string `db:"order_uid"`
}

type Payment struct {
	Transaction  string `db:"transaction" json:"transaction,omitempty"`
	RequestID    string `db:"request_id" json:"request_id,omitempty"`
	Currency     string `db:"currency" json:"currency,omitempty"`
	Provider     string `db:"provider" json:"provider,omitempty"`
	Amount       int    `db:"amount" json:"amount,omitempty"`
	PaymentDt    int    `db:"payment_dt" json:"payment_dt,omitempty"`
	Bank         string `db:"bank" json:"bank,omitempty"`
	DeliveryCost int    `db:"delivery_cost" json:"delivery_cost,omitempty"`
	GoodsTotal   int    `db:"goods_total" json:"goods_total,omitempty"`
	CustomFee    int    `db:"custom_fee" json:"custom_fee,omitempty"`
	OrderUID     string `db:"order_uid"`
}

type Item struct {
	ChrtID      int    `db:"chrt_id" json:"chrt_id,omitempty"`
	TrackNumber string `db:"track_number" json:"track_number,omitempty"`
	Price       int    `db:"price" json:"price,omitempty"`
	Rid         string `db:"rid" json:"rid,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Sale        int    `db:"sale" json:"sale,omitempty"`
	Size        string `db:"size" json:"size,omitempty"`
	TotalPrice  int    `db:"total_price" json:"total_price,omitempty"`
	NmID        int    `db:"nm_id" json:"nm_id,omitempty"`
	Brand       string `db:"brand" json:"brand,omitempty"`
	Status      int    `db:"status" json:"status,omitempty"`
	OrderUID    string `db:"order_uid"`
}
