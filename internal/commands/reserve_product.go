package commands

import (
	"log"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
)

type ReserveProductCommand struct {
	ReserveProduct *models.ReserveProductModel
}

func (cmdProduct *ReserveProductCommand) Execute() error {
	email := cmdProduct.ReserveProduct.CustomerData.Email
	releaseDate := cmdProduct.ReserveProduct.ReleaseDate
	modelName := cmdProduct.ReserveProduct.ModelName

	log.Printf("Reservando o  produto %s com lançamento no dia %s para %s \n", modelName, releaseDate.String(), email)
	return nil
}
