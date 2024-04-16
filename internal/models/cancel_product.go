package models

import "time"

type CancelReserveModel struct {
	ReserveID    string
	CustomerData Customer
}

type HistoryObj struct {
	ReservationID string
	Sku           string
	ReserveDate   *time.Time
}

type Product struct {
	Sku         string
	ReleaseDate *time.Time
	ModelName   string
	Kind        string
}
