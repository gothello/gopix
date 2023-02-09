package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/request"
	"github.com/gothello/go-pix-mercado-pago/service"
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
	service := service.NewServiceMySql(db)

	//create type pix
	//	p := pix.NewPix(0.01, "Pagamento dos Serviços", time.Minute*10, "http://google.com.br", "wpsolucoes@gmail.com")

	//create transaction pix
	// output, err := p.CreatePix()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(output)

	// //save in mysql
	// if err := service.Insert(output); err != nil {
	// 	log.Fatalln(err)
	// }

	// //get by id transaction
	transaction, err := service.GetByIdPayment(54534097538)
	if err != nil {
		log.Fatalln(err)
	}

	//refund payer
	err = transaction.RefundPix()
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Update(transaction)
	if err != nil {
		log.Fatalln(err)
	}

	// //print response json
	// fmt.Printf(transaction.Status)
}
