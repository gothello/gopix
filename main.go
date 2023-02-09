package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	pix "github.com/gothello/go-pix-mercado-pago/create-pix"
	"github.com/gothello/go-pix-mercado-pago/create-pix/service"
	"github.com/gothello/go-pix-mercado-pago/request"
)

func GetMethodsPayments() error {
	h := map[string]string{

		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `Bearer TEST-6812762136376103-020807-45d3ade4692fa87ac9b5b987554b77bf-811772071`,
	}

	opt := request.NewOptions("GET", "https://api.mercadopago.com/v1/payment_methods", "", 3000, h)

	r := opt.Request()
	if r.Err != nil {
		return r.Err
	}

	//fmt.Println(string(r.Body))

	return nil

}

func main() {

	//open db
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.1:3306)/orders")
	if err != nil {
		log.Fatalln(err)
	}

	//return new type service
	//type service.Service implement function insert and get
	service := service.NewService(db)

	//create type pix
	p := pix.NewPix(1, "Pagamento dos Serviços", time.Minute*10, "http://google.com.br", "wpsolucoes@gmail.com")

	//create transaction pix
	output, err := p.CreatePix()
	if err != nil {
		log.Fatalln(err)
	}

	//save in mysql
	if err := service.Insert(output); err != nil {
		log.Fatalln(err)
	}

	//get by id transaction
	idPayment := 54530435785

	transaction, err := service.GetByIdPayment(int64(idPayment))
	if err != nil {
		log.Fatalln(err)
	}

	//refund payer
	err = transaction.RefundPix()
	if err != nil {
		log.Fatalln(err)
	}

	//print response json
	fmt.Printf(transaction.Status)
}
