package vandar

type paymentRequest struct {
	Amount          int    `json:"amount"`
	MobileNumber    string `json:"mobile"`
	Description     string `json:"description"`
	NationalCode    string `json:"national_code"`
	ValidCardNumber string `json:"valid_card_number"`
}

type vandarPaymentRequest struct {
	ApiKey          string `json:"api_key"`
	Amount          int    `json:"amount"`
	CallBackURL     string `json:"callback_url"`
	MobileNumber    string `json:"mobile_number"`
	FactorNumber    string `json:"factorNumber"`
	Description     string `json:"description"`
	NationalCode    string `json:"national_code"`
	ValidCardNumber string `json:"valid_card_number"`
}

type PaymentResponse struct {
	Status int      `json:"status"`
	Token  string   `json:"token"`
	Errors []string `json:"errors"`
}

type verifyRequest struct {
	APIKey string `json:"api_key"`
	Token  string `json:"token"`
}

type vandarVerifyRequest struct {
	APIKey string `json:"api_key"`
	Token  string `json:"token"`
}

type VerifyResponse struct {
	Status       int      `json:"status"`
	Errors       []string `json:"errors"`
	Amount       string   `json:"amount"`
	RealAmount   int      `json:"realAmount"`
	Wage         string   `json:"wage"`
	TransID      int64    `json:"transId"`
	FactorNumber string   `json:"factorNumber"`
	Mobile       string   `json:"mobile"`
	Description  string   `json:"description"`
	CardNumber   string   `json:"cardNumber"`
	PaymentDate  string   `json:"paymentDate"`
	Cid          string   `json:"cid"`
	Message      string   `json:"message"`
}

type transactionDetailRequest struct {
	Token string `json:"token"`
}

type vandarTransactionDetailRequest struct {
	APIKey string `json:"api_key"`
	Token  string `json:"token"`
}

type TransactionDetailResponse struct {
	Status       *int    `json:"status"`
	Amount       *string `json:"amount"`
	TransId      *int    `json:"transId"`
	RefNumber    *string `json:"refnumber"`
	TrackingCode *string `json:"trackingCode"`
	FactorNumber *string `json:"factorNumber"`
	Mobile       *string `json:"mobile"`
	Description  *string `json:"description"`
	CardNumber   *string `json:"cardNumber"`
	CID          *string `json:"CID"`
	CreatedAt    *string `json:"createdAt"`
	PaymentDate  *string `json:"paymentDate"`
	Code         *int    `json:"code"`
}
