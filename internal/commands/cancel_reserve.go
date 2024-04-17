package commands

import (
	"log"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
)

type CancelProductReserveCommand struct {
	CancelProductReverse models.CancelReserveModel
}

func (cmdCancel *CancelProductReserveCommand) Execute() error {
	log.Printf("Id da Reserva: %s \n", cmdCancel.CancelProductReverse.ReserveID)
	return nil
}
