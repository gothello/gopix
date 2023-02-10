package usecase

import (
	"time"

	"github.com/google/uuid"
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/service"
)

type InputPix struct {
	ID                string  `json:"id"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"descryption"`
	ExpirationMinutes int     `json:"expiration_minute"`
	UrlNotify         string  `json:"url_notify"`
	Email             string  `json:"email"`
}

type CreatePixUseCase struct {
	PixRepositoryUseCase service.PixRepositoryUseCase
}

func NewCreatePixUseCase(s service.PixRepositoryUseCase) *CreatePixUseCase {
	return &CreatePixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s CreatePixUseCase) Execute(input InputPix) (*pix.OutputPix, error) {

	if input.ExpirationMinutes == 0 {
		input.ExpirationMinutes = 10
	}

	i := &pix.InputPix{
		ID:               uuid.New().String(),
		Amount:           input.Amount,
		Description:      input.Description,
		TimeOfExpiration: time.Minute * time.Duration(input.ExpirationMinutes),
		UrlNotify:        input.UrlNotify,
		Email:            input.Email,
	}

	output, err := i.CreatePix()

	if err != nil {
		return nil, err
	}

	if err := s.PixRepositoryUseCase.Insert(output); err != nil {
		return nil, err
	}

	return output, nil
}
