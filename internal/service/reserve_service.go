package service

import (
	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/commands"
	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
	"github.com/Israel-Ferreira/pre-order/preorder-service/pkg/requests"
)

func SendReserve(preOrderReq requests.ReserveProductReq, product models.Product) error {
	rspm := models.NewReserveProductModel(&preOrderReq, product)

	cmdReserve := commands.ReserveProductCommand{
		ReserveProduct: rspm,
	}

	if err := ExecuteCommand(&cmdReserve); err != nil {
		return err
	}

	SaveInHistory(*rspm)

	return nil
}

func CancelReserve(historyObj models.HistoryObj) error {
	cmdCancel := commands.CancelProductReserveCommand{
		CancelProductReverse: models.CancelReserveModel{
			ReserveID: historyObj.ReservationID,
			Sku:       historyObj.Sku,
			CustomerData: models.Customer{
				Email: historyObj.CustomerData.Email,
				Cpf:   historyObj.CustomerData.Cpf,
			},
		},
	}

	if err := ExecuteCommand(&cmdCancel); err != nil {
		return err
	}

	return nil
}
