package vandar

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func getTransactionDetail(token string) (*TransactionDetailResponse, error) {

	postBody, _ := json.Marshal(vandarTransactionDetailRequest{
		APIKey: ApiKey,
		Token:  token,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://ipg.vandar.io/api/v3/transaction", "application/json", responseBody)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var detail *TransactionDetailResponse
	err = json.NewDecoder(resp.Body).Decode(&detail)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func TransactionDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	var request transactionDetailRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	detail, err := getTransactionDetail(request.Token)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	jData, err := json.Marshal(detail)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	w.Write(jData)

	return

}
