package service

import (
	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
	"github.com/Israel-Ferreira/pre-order/preorder-service/pkg/interfaces"
)

func SaveInHistory(obj models.ReserveProductModel) (models.HistoryObj, error) {
	return models.HistoryObj{}, nil
}

func ExecuteCommand(cmd interfaces.Command) (errFunc error) {
	if err := cmd.Execute(); err != nil {
		errFunc = err
	}

	return
}
