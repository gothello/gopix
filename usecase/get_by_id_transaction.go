package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type GetByIdPixUseCase struct {
	ServiceUseCase service.PixRepositoryUseCase
}

func NewGetByIdPixUseCase(s service.PixRepositoryUseCase) *GetByIdPixUseCase {
	return &GetByIdPixUseCase{
		ServiceUseCase: s,
	}
}

func (s GetByIdPixUseCase) Execute(paymentID int64) (*pix.OutputPix, error) {
	output, err := s.ServiceUseCase.GetByIdPayment(paymentID)
	if err != nil {
		return nil, err
	}

	return output, nil
}
