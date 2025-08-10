package vo

type RoomSelected struct {
	ID       string `json:"id"`
	Quantity uint8  `json:"quantity"`
}

type CreatePaymentURLInput struct {
	CheckIn         string         `json:"check_in"`
	CheckOut        string         `json:"check_out"`
	AccommodationID string         `json:"accommodation_id"`
	RoomSelected    []RoomSelected `json:"room_selected"`
}
type CreatePaymentURLOutput struct {
	Url string `json:"url"`
}

type VNPayParams map[string]string

type VNPayReturnInput struct {
}

type VNPayIPNInput struct {
}

type VNPayResponse struct {
	RspCode string `json:"response_code"`
	Message string `json:"message"`
}

type PostQueryDRInput struct {
	OrderID   string `json:"orderId" form:"orderId"`
	TransDate string `json:"transDate" form:"transDate"`
}

type QueryDataObj struct {
	VnpRequestID       string `json:"vnp_RequestId"`
	VnpVersion         string `json:"vnp_Version"`
	VnpCommand         string `json:"vnp_Command"`
	VnpTmnCode         string `json:"vnp_TmnCode"`
	VnpTxnRef          string `json:"vnp_TxnRef"`
	VnpOrderInfo       string `json:"vnp_OrderInfo"`
	VnpTransactionDate string `json:"vnp_TransactionDate"`
	VnpCreateDate      string `json:"vnp_CreateDate"`
	VnpIpAddr          string `json:"vnp_IpAddr"`
	VnpSecureHash      string `json:"vnp_SecureHash"`
}

type PostRefundInput struct {
	OrderID   string `json:"orderId" form:"orderId"`
	TransDate string `json:"transDate" form:"transDate"`
	Amount    string `json:"amount" form:"amount"`
	TransType string `json:"transType" form:"transType"`
	User      string `json:"user" form:"user"`
}

type RefundDataObj struct {
	VnpRequestID       string `json:"vnp_RequestId"`
	VnpVersion         string `json:"vnp_Version"`
	VnpCommand         string `json:"vnp_Command"`
	VnpTmnCode         string `json:"vnp_TmnCode"`
	VnpTransactionType string `json:"vnp_TransactionType"`
	VnpTxnRef          string `json:"vnp_TxnRef"`
	VnpAmount          string `json:"vnp_Amount"`
	VnpTransactionNo   string `json:"vnp_TransactionNo"`
	VnpCreateBy        string `json:"vnp_CreateBy"`
	VnpOrderInfo       string `json:"vnp_OrderInfo"`
	VnpTransactionDate string `json:"vnp_TransactionDate"`
	VnpCreateDate      string `json:"vnp_CreateDate"`
	VnpIpAddr          string `json:"vnp_IpAddr"`
	VnpSecureHash      string `json:"vnp_SecureHash"`
}

type PaymentResultData struct {
	OrderIDExternal string `json:"order_id"`
	ResponseCode    string `json:"response_code"`
	Amount          string `json:"amount"`
	BankCode        string `json:"bank_code"`
	TransactionNo   string `json:"transaction_no"`
	PayDate         string `json:"pay_date"`
}
