package usecase

import (
	"github.com/gothello/go-pix-mercado-pago/entity"
)

type OneOutputPix struct {
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
type FindPixUseCase struct {
	PixRepositoryUseCase entity.PixRepository
}

func NewFindPixUseCase(s entity.PixRepository) *FindPixUseCase {
	return &FindPixUseCase{
		PixRepositoryUseCase: s,
	}
}

func (s FindPixUseCase) Execute(paymentID int64) (*OneOutputPix, error) {
	output, err := s.PixRepositoryUseCase.GetByIdPayment(paymentID)
	if err != nil {
		return nil, err
	}

	//	fmt.Println(output.Ticket)

	return &OneOutputPix{
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
	}, nil
}
