package usecase

import (
	"github.com/gothello/go-pix-mercado-pago/entity"
)

type OutputCancel struct {
	ID     string  `json:"id"`
	IdPay  int64   `json:"idpay"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
	Email  string  `json:"email"`
}

type CancelUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewCancelUseCase(s entity.PixRepository) *CancelUseCase {
	return &CancelUseCase{
		PixRepositoryUseCase: s,
	}
}

func (u CancelUseCase) Execute(IDpay int64) (*OutputCancel, error) {
	o, err := u.PixRepositoryUseCase.GetByIdPayment(IDpay)
	if err != nil {
		return nil, err
	}

	err = o.CancelPix()
	if err != nil {
		return nil, err
	}

	if err := u.PixRepositoryUseCase.Update(o); err != nil {
		return nil, err
	}

	return &OutputCancel{
		ID:     o.ID,
		IdPay:  o.IDExternalTransaction,
		Status: o.Status,
		Amount: o.Amount,
		Email:  o.Email,
	}, nil
}
