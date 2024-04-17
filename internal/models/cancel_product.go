package models

import "time"

type CancelReserveModel struct {
	ReserveID    string
	Sku          string
	CustomerData Customer
}

type HistoryObj struct {
	ReservationID string
	Sku           string
	ReserveDate   *time.Time
	CustomerData  Customer
}

type Product struct {
	Sku         string
	ReleaseDate *time.Time
	ModelName   string
	Kind        string
}
