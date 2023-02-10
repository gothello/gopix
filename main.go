package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/service"
	"github.com/gothello/go-pix-mercado-pago/usecase"
)

func main() {

	//open db
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.1:3306)/orders")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	service := service.NewServiceMySql(db)
	//	create := usecase.NewCreatePixUseCase(service)
	//	findAll := usecase.NewFindAllPixUseCase(service)
	getByIdPay := usecase.NewGetByIdPixUseCase(service)
	// update := usecase.NewUpdatePixUseCase(service)
	// cancel := usecase.NewCancelUseCase(service)
	// refund := usecase.NewRefundUseCase(service)

	// h := web.PixHandlers(create, cancel, refund, getByIdPay, update, cancel, refund)

	//			    Amount      Descryption         time expiration   webhook notify pay     email user payed
	// p := pix.NewPix(0.01, "Pagamento dos Serviços", time.Minute*10, "http://google.com.br", "wpsolucoes@gmail.com")

	// out, err := create.Execute(p)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	al, err := getByIdPay.Execute(54548671739)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(al)

	// fmt.Printf("%#v\n", out)
}
