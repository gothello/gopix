package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/utils"
	"github.com/gothello/go-pix-mercado-pago/web"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	QUEUES = map[string]string{
		"CREATE":   "newpix",
		"CREATED":  "createdpix",
		"APPROVED": "approved",
	}

	API_PORT        = os.Getenv("API_PORT")
	SECRET_AUTH_KEY = os.Getenv("SECRET_AUTH_KEY")
)

func init() {

	if SECRET_AUTH_KEY == "" {
		var err error

		SECRET_AUTH_KEY, err = utils.GenerateAuthAcess()
		if err != nil {
			log.Fatalln(err)
		}

		utils.SECRET_AUTH_KEY = SECRET_AUTH_KEY

		color.Red("YOU NOT PREDEFINED KEY")
		color.Green("SYSTEM GENERATE KEY FOR YOU")
		color.Green("KEY AUTH SAVED ON /HOME %s", SECRET_AUTH_KEY)
	}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	conn, err := amqp.Dial("amqp://admin:admin@rabbit:5672/")
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	rep := entity.NewRespositoryMySql(db)
	handlers := web.LoadAllUseCases(rep)
	handlers.LoadRoutes()

	rabbitUseCase := usecase.NewRabbitConnectionUseCase(conn, rep)

	input := make(chan rabbit.RabbitInputChan)
	go rabbit.Consumer(conn, QUEUES["CREATE"], input)
	go rabbitUseCase.RabbitCreatePixUseCase(input, rep, QUEUES)

	log.Printf("api running port %s\n", API_PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", API_PORT), nil); err != nil {
		log.Fatalln(err)
	}
}
