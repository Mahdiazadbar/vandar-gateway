package vandar

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

var persianMobileValidatorRegexp = regexp.MustCompile("^((\\+98|0)9\\d{9})$")

func getVandarToken(amount int, mobile, nationalCode, validCardNumber, description string) (*PaymentResponse, error) {

	var request = vandarPaymentRequest{
		ApiKey:          ApiKey,
		CallBackURL:     CallbackUrl,
		Amount:          amount,
		MobileNumber:    mobile,
		NationalCode:    nationalCode,
		ValidCardNumber: validCardNumber,
		Description:     description,
	}
	postBody, _ := json.Marshal(request)

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://ipg.vandar.io/api/v3/send", "application/json", responseBody)

	if err != nil {
		return nil, err
	}

	var paymentResponse *PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)

	if err != nil {
		return nil, err
	}

	return paymentResponse, nil

}

func GetToken(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request paymentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mobile := request.MobileNumber
	if !persianMobileValidatorRegexp.MatchString(mobile) {
		http.Error(w, "Mobile is invalid", 422)
		return
	}

	if request.Amount < 1 {
		http.Error(w, "amount is invalid", 422)
		return
	}

	if request.ValidCardNumber != "" && len(request.ValidCardNumber) != 16 {
		http.Error(w, "Valid Card Number  is invalid", 422)
		return
	}

	if request.NationalCode != "" && len(request.NationalCode) != 10 {
		http.Error(w, "National Code  is invalid", 422)
		return
	}

	paymentResponse, err := getVandarToken(request.Amount, request.MobileNumber, request.NationalCode, request.ValidCardNumber, request.Description)
	jData, err := json.Marshal(paymentResponse)

	w.Write(jData)
	return
}
