package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/service"
	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/web"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.1:3306)/orders")
	if err != nil {
		log.Fatalln(err)
	}

	service := service.NewServiceMySql(db)

	create := usecase.NewCreatePixUseCase(service)
	cancel := usecase.NewCancelUseCase(service)
	refund := usecase.NewRefundUseCase(service)
	find := usecase.NewFindPixUseCase(service)
	findall := usecase.NewFindAllPixUseCase(service)

	h := web.NewPixHandlers(create, cancel, refund, find, findall)

	http.HandleFunc("/create", h.Create)
	http.HandleFunc("/cancel", h.Cancel)
	http.HandleFunc("/refund", h.Refund)
	http.HandleFunc("/find", h.Find)
	http.HandleFunc("/all", h.FindAll)

	log.Println("api running port 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln(err)
	}

	// out, err := create.Execute(p)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// al, err := getByIdPay.Execute(54548671739)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//fmt.Println(al)

	//fmt.Printf("%#v\n", out)
}
