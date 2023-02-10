package usecase

import (
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type CreatePixUseCase struct {
	ServiceUseCase service.PixRepositoryUseCase
}

func NewCreatePixUseCase(s service.PixRepositoryUseCase) *CreatePixUseCase {
	return &CreatePixUseCase{
		ServiceUseCase: s,
	}
}

func (s CreatePixUseCase) Execute(input *pix.InputPix) (*pix.OutputPix, error) {
	output, err := input.CreatePix()
	if err != nil {
		return nil, err
	}

	if err := s.ServiceUseCase.Insert(output); err != nil {
		return nil, err
	}

	return output, nil
}
