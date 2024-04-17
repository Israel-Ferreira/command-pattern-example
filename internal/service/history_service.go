package service

import (
	"errors"
	"time"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
)

var history []models.HistoryObj

func SaveInHistory(obj models.ReserveProductModel) {

	reserveDate := time.Now()

	historyObj := models.HistoryObj{
		ReservationID: obj.ReservationID,
		Sku:           obj.Sku,
		ReserveDate:   &reserveDate,
		CustomerData: models.Customer{
			Email: obj.CustomerData.Email,
			Cpf:   obj.CustomerData.Cpf,
		},
	}

	history = append(history, historyObj)
}

func FindHistory(reservationId string) (models.HistoryObj, error) {
	var historyObj models.HistoryObj

	for _, obj := range history {
		if obj.ReservationID == reservationId {
			historyObj = obj
			break
		}
	}

	if historyObj.ReservationID == "" {
		return historyObj, errors.New("erro ao encontrar o id da reserva no historico")
	}

	return historyObj, nil
}
