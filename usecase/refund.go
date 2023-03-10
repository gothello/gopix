package usecase

import "github.com/gothello/go-pix-mercado-pago/entity"

type OutputRefund struct {
	ID     string  `json:"id"`
	IdPay  int64   `json:"idpay"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
	Email  string  `json:"email"`
}

type RefundUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewRefundUseCase(s entity.PixRepository) *RefundUseCase {
	return &RefundUseCase{
		PixRepositoryUseCase: s,
	}
}

func (u *RefundUseCase) Execute(IDpay int64) (*OutputRefund, error) {
	o, err := u.PixRepositoryUseCase.GetByIdPayment(IDpay)
	if err != nil {
		return nil, err
	}

	if err := o.RefundPix(); err != nil {
		return nil, err
	}

	if err := u.PixRepositoryUseCase.Update(o); err != nil {
		return nil, err
	}

	return &OutputRefund{
		o.ID,
		o.IDExternalTransaction,
		o.Status,
		o.Amount,
		o.Email,
	}, nil
}
