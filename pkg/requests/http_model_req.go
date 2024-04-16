package requests

type ReserveProductReq struct {
	Sku           string `json:"sku"`
	CustomerName  string `json:"customerName"`
	CustomerEmail string `json:"customerEmail"`
	CustomerCpf   string `json:"customerCpf"`
}
