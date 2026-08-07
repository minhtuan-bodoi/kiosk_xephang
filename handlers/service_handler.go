package handlers

import "net/http"

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("service handler placeholder"))
}
