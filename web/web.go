package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/utils"
)

var (
	IDNotExist = "id not exist"
)

type PixHandlers struct {
	CreatePixUseCase  *usecase.CreatePixUseCase
	CancelPixUseCase  *usecase.CancelUseCase
	RefundPixUseCase  *usecase.RefundUseCase
	FindPixUsecase    *usecase.FindPixUseCase
	FindAllPixUseCase *usecase.FindAllPixUseCase
}

func NewPixHandlers(create *usecase.CreatePixUseCase, cancel *usecase.CancelUseCase, refund *usecase.RefundUseCase, find *usecase.FindPixUseCase, findall *usecase.FindAllPixUseCase) *PixHandlers {
	return &PixHandlers{
		CreatePixUseCase:  create,
		CancelPixUseCase:  cancel,
		RefundPixUseCase:  refund,
		FindPixUsecase:    find,
		FindAllPixUseCase: findall,
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
