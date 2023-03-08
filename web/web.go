package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/utils"
)

var (
	IDNotExist = "id not exist"
	EmailEmpty = "email is empty"
	Empty      = ""
)

type PixHandlers struct {
	CreatePixUseCase  *usecase.CreatePixUseCase
	CancelPixUseCase  *usecase.CancelUseCase
	RefundPixUseCase  *usecase.RefundUseCase
	FindPixUsecase    *usecase.FindPixUseCase
	FindAllPixUseCase *usecase.FindAllPixUseCase
	NotifyPixUseCase  *usecase.NotifyPixUseCase
}

func NewPixHandlers(create *usecase.CreatePixUseCase, cancel *usecase.CancelUseCase, refund *usecase.RefundUseCase, find *usecase.FindPixUseCase, findall *usecase.FindAllPixUseCase, notify *usecase.NotifyPixUseCase) *PixHandlers {
	return &PixHandlers{
		CreatePixUseCase:  create,
		CancelPixUseCase:  cancel,
		RefundPixUseCase:  refund,
		FindPixUsecase:    find,
		FindAllPixUseCase: findall,
		NotifyPixUseCase:  notify,
	}
}

func (h *PixHandlers) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ToErro(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var input usecase.InputPix

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		utils.ToErro(w, http.StatusText(http.StatusFailedDependency), http.StatusFailedDependency)
		return
	}

	if input.Email == Empty {
		utils.ToErro(w, EmailEmpty, http.StatusBadRequest)
		return
	}

	out, err := h.CreatePixUseCase.Execute(input)
	if err != nil {
		log.Println(err)
		utils.ToErro(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.ToJson(w, out, http.StatusCreated)
}

func (h *PixHandlers) Cancel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ToErro(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := utils.ParseID(r)
	if err != nil {
		utils.ToErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.CancelPixUseCase.Execute(id)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			utils.ToErro(w, IDNotExist, http.StatusBadRequest)
			return
		}

		utils.ToErro(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.ToJson(w, out, http.StatusOK)
}

func (h *PixHandlers) Refund(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ToErro(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := utils.ParseID(r)
	if err != nil {
		utils.ToErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.RefundPixUseCase.Execute(id)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			utils.ToErro(w, IDNotExist, http.StatusBadRequest)
			return
		}

		utils.ToErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ToJson(w, out, http.StatusOK)
}

func (h *PixHandlers) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ToErro(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	all, err := h.FindAllPixUseCase.Execute()
	if err != nil {
		log.Println(err)
		utils.ToErro(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.ToJson(w, all, http.StatusOK)
}

func (h *PixHandlers) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ToErro(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := utils.ParseID(r)
	if err != nil {
		utils.ToErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.FindPixUsecase.Execute(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ToErro(w, IDNotExist, http.StatusBadRequest)
			return
		}
		utils.ToErro(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.ToJson(w, out, http.StatusOK)
}

func (h *PixHandlers) WebHook(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		fmt.Println(err)
		return
	}

	action, ok := response["action"].(string)
	if !ok {
		return
	}

	if action == "payment.updated" {
		data, ok := response["data"].(map[string]interface{})
		if !ok {
			return
		}

		idpay, ok := data["id"]
		if !ok {
			return
		}

		go func() {

			if err := h.NotifyPixUseCase.Execute(idpay.(string)); err != nil {
				log.Println(err)
			}
		}()
	}

	w.WriteHeader(http.StatusOK)
}
