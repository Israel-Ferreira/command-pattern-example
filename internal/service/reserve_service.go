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

	return nil
}
