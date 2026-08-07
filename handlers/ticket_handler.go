package handlers

import "net/http"

func TicketHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ticket handler placeholder"))
}
