package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/pix"
)

type InputPix struct {
	ID                string  `json:"id"`
	User              string  `json:"user"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"descryption"`
	ExpirationMinutes int     `json:"expiration_minutes"`
	UrlNotify         string  `json:"url_notify"`
	Email             string  `json:"email"`
}

type OutputPix struct {
	ID           string  `json:"id"`
	IDPAY        int64   `json:"id_pay"`
	CreateAt     string  `json:"created_at"`
	ExpiresAt    string  `json:"expires_at"`
	Status       string  `json:"status"`
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	Email        string  `json:"email"`
	QrCode       string  `json:"qrcode"`
	QrCodeBase64 string  `json:"qrcodebase"`
}

// createdpixusecase conteins atribute pixrepository usecase
type CreatePixUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewCreatePixUseCase(s entity.PixRepository) *CreatePixUseCase {
	return &CreatePixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s CreatePixUseCase) Execute(input InputPix) (*OutputPix, error) {
	//set default time expiration if user no define time
	if input.ExpirationMinutes < 10 {
		input.ExpirationMinutes = 10
	}

	if input.Email == "" {
		return nil, errors.New("email not informed")
	}

	if input.UrlNotify == "" {
		input.UrlNotify = "https://fea6-170-245-238-18.sa.ngrok.io/notify"
	}

	//tranform type input pacakge usecase into *pix.InputPix
	i := &pix.InputPix{
		ID:               uuid.New().String(),
		Amount:           input.Amount,
		Description:      input.Description,
		TimeOfExpiration: time.Duration(input.ExpirationMinutes) * time.Minute,
		UrlNotify:        input.UrlNotify,
		Email:            input.Email,
	}

	//send request to api mercadopago
	output, err := i.CreatePix()
	if err != nil {
		return nil, err
	}

	//save in database sql
	if err := s.PixRepositoryUseCase.Insert(output); err != nil {
		return nil, err
	}

	//return data information of transaction
	return &OutputPix{
		ID:           output.ID,
		IDPAY:        output.IDExternalTransaction,
		CreateAt:     output.CreateAt,
		ExpiresAt:    output.ExpiresAt,
		Status:       output.Status,
		Type:         output.Type,
		Amount:       output.Amount,
		Email:        input.Email,
		QrCode:       output.QrCode,
		QrCodeBase64: output.QrCodeBase64,
	}, nil
}
