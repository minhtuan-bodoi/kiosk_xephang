package handlers

import "net/http"

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("counter handler placeholder"))
}
