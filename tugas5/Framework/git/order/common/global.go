package common

type Message struct {
	Code   int     `json:"code"`
	Remark string  `json:"remark"`
	Result *Result `json:"result,omitempty"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type StatusRequest struct {
	Request    string `json:"request"`
	TrxID      string `json:"trx_id"`
	MerchantID string `json:"merchant_id"`
	BillNO     string `json:"bill_no"`
}

type StatusResponse struct {
	Response          string `json:"response"`
	TrxID             string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNO            string `json:"bill_no"`
	PaymentReff       string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}
