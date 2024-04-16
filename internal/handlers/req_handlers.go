package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/service"
	"github.com/Israel-Ferreira/pre-order/preorder-service/pkg/requests"
)

func sendErrorMsg(rw http.ResponseWriter, err error, statusCode int) {
	errorMsg := map[string]any{
		"msg":        err.Error(),
		"statusCode": statusCode,
	}

	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(errorMsg)
}

func ReservePreOrderProductHandler(rw http.ResponseWriter, r *http.Request) {

	var preOrderReq requests.ReserveProductReq

	if err := json.NewDecoder(r.Body).Decode(&preOrderReq); err != nil {
		sendErrorMsg(rw, err, http.StatusBadRequest)
		return
	}

	product, err := service.FindProductBySKU(preOrderReq.Sku)

	if err != nil {
		sendErrorMsg(rw, err, http.StatusNotFound)
		return
	}

	if err = service.SendReserve(preOrderReq, product); err != nil {
		sendErrorMsg(rw, err, http.StatusInternalServerError)
		return
	}

	msg := map[string]any{
		"msg": "Reserva Realizada com Sucesso",
	}

	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(msg); err != nil {
		sendErrorMsg(rw, err, http.StatusInternalServerError)
		return
	}

}
