package usecase

import "github.com/gothello/go-pix-mercado-pago/service"

type OutputCancel struct {
	ID     string  `json:"id"`
	IdPay  int64   `json:"idpay"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type CancelUseCase struct {
	PixRepositoryUseCase service.PixRepositoryUseCase
}

func NewCancelUseCase(s service.PixRepositoryUseCase) *CancelUseCase {
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
	}, nil
}
