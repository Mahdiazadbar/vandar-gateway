package vandar

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	token := r.URL.Query().Get("token")
	http.Redirect(w, r, "https://ipg.vandar.io/v3/"+token, http.StatusSeeOther)
}
