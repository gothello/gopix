package web

import (
	"net/http"

	"github.com/gothello/go-pix-mercado-pago/utils"
)

func (h *PixHandlers) LoadRoutes() {
	http.HandleFunc("/create", utils.IsAdmin(h.Create))
	http.HandleFunc("/cancel", utils.IsAdmin(h.Cancel))
	http.HandleFunc("/refund", utils.IsAdmin(h.Refund))
	http.HandleFunc("/find", utils.IsAdmin(h.Find))
	http.HandleFunc("/all", utils.IsAdmin(h.FindAll))
}
