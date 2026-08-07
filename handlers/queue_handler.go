package handlers

import "net/http"

func QueueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("queue handler placeholder"))
}
