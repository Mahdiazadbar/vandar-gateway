package vandar

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func verifyTransaction(token string) (*VerifyResponse, error) {

	postBody, _ := json.Marshal(vandarVerifyRequest{
		APIKey: ApiKey,
		Token:  token,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://ipg.vandar.io/api/v3/verify", "application/json", responseBody)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var verify *VerifyResponse
	err = json.NewDecoder(resp.Body).Decode(&verify)
	if err != nil {
		return nil, err
	}

	return verify, nil
}

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	var request verifyRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	verify, err := verifyTransaction(request.Token)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	jData, err := json.Marshal(verify)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	w.Write(jData)

	return
}
