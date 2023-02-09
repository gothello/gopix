package service

import (
	"database/sql"

	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	//pix "github.com/gothello/go-pix-mercado-pago/create-pix"
)

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

type ServiceUseCase interface {
	Insert(*pix.OutputPix)
	GetByIdPayment(int64)
}

func (s *Service) Insert(p *pix.OutputPix) error {

	_, err := s.DB.Exec("INSERT into pix4(id, id_payment, created_at, expires_at, type, amount, ticket, qrcode, qrcodebase) values (?,?,?,?,?,?,?,?,?)", p.ID, p.IDExternalTransaction, p.CreateAt, p.ExpiresAt, p.Type, p.Amount, p.Ticket, p.QrCode, p.QrCodeBase64)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByIdPayment(id int64) (*pix.OutputPix, error) {
	rows, err := s.DB.Prepare("select id, id_payment, created_at, expires_at ,type, amount, ticket, qrcode, qrcodebase from pix4 where id_payment = ?")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var p pix.OutputPix

	err = rows.QueryRow(id).Scan(&p.ID, &p.IDExternalTransaction, &p.CreateAt, &p.ExpiresAt, &p.Type, &p.Amount, &p.Ticket, &p.QrCode, &p.QrCodeBase64)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
