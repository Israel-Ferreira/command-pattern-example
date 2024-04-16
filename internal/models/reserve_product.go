package models

import (
	"time"

	"github.com/Israel-Ferreira/pre-order/preorder-service/pkg/requests"
	"github.com/google/uuid"
)

type Customer struct {
	Name  string
	Email string
	Cpf   string
}

type ReserveProductModel struct {
	ReservationID string
	Sku           string
	ModelName     string
	ReleaseDate   *time.Time
	Kind          string
	CustomerData  *Customer
}

func NewReserveProductModel(r *requests.ReserveProductReq, p Product) *ReserveProductModel {
	reservationID, _ := uuid.NewUUID()

	return &ReserveProductModel{
		ReservationID: reservationID.String(),
		Sku:           p.Sku,
		ModelName:     p.ModelName,
		ReleaseDate:   p.ReleaseDate,
		Kind:          p.Kind,
		CustomerData: &Customer{
			Name:  r.CustomerName,
			Email: r.CustomerEmail,
			Cpf:   r.CustomerCpf,
		},
	}

}
