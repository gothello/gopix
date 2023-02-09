package service

import (
	"database/sql"

	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	//pix "github.com/gothello/go-pix-mercado-pago/create-pix"
)

type ServiceMySql struct {
	DB *sql.DB
}

func NewServiceMySql(db *sql.DB) *ServiceMySql {
	return &ServiceMySql{
		DB: db,
	}
}

type ServiceUseCase interface {
	Insert(*pix.OutputPix)
	GetByIdPayment(int64) (*pix.OutputPix, error)
	GetAll() ([]*pix.OutputPix, error)
	Update(*pix.OutputPix) error
}

func (s *ServiceMySql) Insert(p *pix.OutputPix) error {

	_, err := s.DB.Exec("INSERT into pix4(id, id_payment, created_at, expires_at, status, type, amount, ticket, qrcode, qrcodebase) values (?,?,?,?,?,?,?,?,?)", p.ID, p.IDExternalTransaction, p.CreateAt, p.ExpiresAt, p.Status, p.Type, p.Amount, p.Ticket, p.QrCode, p.QrCodeBase64)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceMySql) GetByIdPayment(id int64) (*pix.OutputPix, error) {
	rows, err := s.DB.Prepare("select id, id_payment, created_at, expires_at ,status, type, amount, ticket, qrcode, qrcodebase from pix4 where id_payment = ?")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var p pix.OutputPix

	err = rows.QueryRow(id).Scan(&p.ID, &p.IDExternalTransaction, &p.CreateAt, &p.ExpiresAt, &p.Status, &p.Type, &p.Amount, &p.Ticket, &p.QrCode, &p.QrCodeBase64)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *ServiceMySql) GetAll() ([]*pix.OutputPix, error) {
	rows, err := s.DB.Query("select * from pix4")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trs []*pix.OutputPix

	for rows.Next() {
		var p pix.OutputPix

		if err = rows.Scan(&p.ID, &p.IDExternalTransaction, &p.CreateAt, &p.ExpiresAt, &p.Status, &p.Type, &p.Amount, &p.Ticket, &p.QrCode, &p.QrCodeBase64); err != nil {
			return nil, err
		}

		trs = append(trs, &p)
	}

	return trs, nil
}

func (s *ServiceMySql) Update(p *pix.OutputPix) error {
	stmt, err := s.DB.Prepare("update pix4 set id=?, id_payment=?, created_at=?, expires_at=? ,status=?, type=?, amount=?, ticket=?, qrcode=?, qrcodebase=? from pix4 where id=? ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&p.ID, &p.IDExternalTransaction, &p.CreateAt, &p.ExpiresAt, &p.Status, &p.Type, &p.Amount, &p.Ticket, &p.QrCode, &p.QrCodeBase64, p.ID)
	if err != nil {
		return err
	}

	return nil
}