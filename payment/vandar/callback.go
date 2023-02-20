package vandar

import (
	"encoding/json"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	status := r.URL.Query().Get("payment_status")
	if status != "OK" {
		http.Error(w, "status is not ok", 400)
		return
	}

	token := r.URL.Query().Get("token")

	detail, err := getTransactionDetail(token)

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
