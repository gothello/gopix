package usecase

import (
	"github.com/gothello/go-pix-mercado-pago/entity"
)

type AllOutputPix struct {
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

type FindAllPixUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewFindAllPixUseCase(s entity.PixRepository) *FindAllPixUseCase {
	return &FindAllPixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s FindAllPixUseCase) Execute() ([]*AllOutputPix, error) {
	outputs, err := s.PixRepositoryUseCase.GetAll()
	if err != nil {
		return nil, err
	}

	var all []*AllOutputPix

	for _, output := range outputs {
		all = append(all, &AllOutputPix{
			ID:           output.ID,
			IDPAY:        output.IDExternalTransaction,
			CreateAt:     output.CreateAt,
			ExpiresAt:    output.ExpiresAt,
			Status:       output.Status,
			Type:         output.Type,
			Amount:       output.Amount,
			Email:        output.Email,
			QrCode:       output.QrCode,
			QrCodeBase64: output.QrCodeBase64,
		})
	}

	return all, nil
}
