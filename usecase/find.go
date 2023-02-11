package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type FindPixUseCase struct {
	PixRepositoryUseCase service.PixRepositoryUseCase
}

func NewFindPixUseCase(s service.PixRepositoryUseCase) *FindPixUseCase {
	return &FindPixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s FindPixUseCase) Execute(paymentID int64) (*pix.OutputPix, error) {
	output, err := s.PixRepositoryUseCase.GetByIdPayment(paymentID)
	if err != nil {
		return nil, err
	}

	return output, nil
}
