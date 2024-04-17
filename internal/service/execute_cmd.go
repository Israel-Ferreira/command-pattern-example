package service

import (
	"github.com/Israel-Ferreira/pre-order/preorder-service/pkg/interfaces"
)

func ExecuteCommand(cmd interfaces.Command) (errFunc error) {
	if err := cmd.Execute(); err != nil {
		errFunc = err
	}

	return
}
