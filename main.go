package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gothello/go-pix-mercado-pago/config"
	"github.com/gothello/go-pix-mercado-pago/entity"
	"github.com/gothello/go-pix-mercado-pago/pix"
	"github.com/gothello/go-pix-mercado-pago/rabbit"
	"github.com/gothello/go-pix-mercado-pago/usecase"
	"github.com/gothello/go-pix-mercado-pago/utils"
	"github.com/gothello/go-pix-mercado-pago/web"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
)

var (
	API_PORT = 0

	RABBITMQ = map[string]string{}
	QUEUES   = map[string]string{}
	MYSQL    = map[string]string{}

	BASE_URL_RABBITMQ = "amqp://%s:%s@%s:%s/"
	BASE_URL_MYSQL    = "%s:%s@tcp(%s:%s)/%s"
)

func init() {
	conf, err := config.LoadConfig()
	if err != nil {
		color.Red("ERROR: ", err)
		os.Exit(0)
	}

	if conf.GetString("SECRET_KEY_API") == "" {
		color.Red("ERROR: you not define you secret key to permit access on api")
		os.Exit(0)
	}

	utils.SECRET_AUTH_KEY = conf.GetString("SECRET_KEY_API")

	if conf.GetString("SECRET_MERCADO_PAGO") == "" {
		color.Red("ERROR: you not define secret key to permit access on api MERCADO PAGO")
		os.Exit(0)
	}

	pix.SECRET_KEY_MP = conf.GetString("SECRET_MERCADO_PAGO")

	MYSQL = conf.GetStringMapString("MYSQL")
	RABBITMQ = conf.GetStringMapString("RABBITMQ")
	QUEUES = conf.GetStringMapString("QUEUES")

	if conf.GetInt("API_PORT") != 0 {
		API_PORT = conf.GetInt("API_PORT")
		return
	}

	API_PORT = 4000

}

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf(BASE_URL_MYSQL, MYSQL["user"], MYSQL["pass"], MYSQL["host"], MYSQL["port"], MYSQL["database"]))
	if err != nil {
		log.Fatalf("ERRO: open connection mysql: %v\n", err)
	}

	defer db.Close()

	conn, err := amqp.Dial(fmt.Sprintf(BASE_URL_RABBITMQ, RABBITMQ["user"], RABBITMQ["pass"], RABBITMQ["host"], RABBITMQ["port"]))
	if err != nil {
		log.Fatalf("ERRO: open connection rabbitmq: %v\n", err)
	}

	defer conn.Close()

	rep := entity.NewRespositoryMySql(db)
	handlers := web.LoadAllUseCases(rep)
	handlers.LoadRoutes()

	rabbitUseCase := usecase.NewRabbitConnectionUseCase(conn, rep)
	usecase.NETIFY_CONNECTION = conn

	icreate := make(chan rabbit.RabbitInputChan)
	icancel := make(chan rabbit.RabbitInputChan)
	irefund := make(chan rabbit.RabbitInputChan)

	go rabbit.Consumer(conn, QUEUES["create"], icreate)
	go rabbit.Consumer(conn, QUEUES["cancel"], icancel)
	go rabbit.Consumer(conn, QUEUES["refund"], irefund)

	go rabbitUseCase.RabbitCreatePixUseCase(icreate, rep, QUEUES)
	go rabbitUseCase.RabbitCancelPixUseCase(icancel, rep, QUEUES)
	go rabbitUseCase.RabbitRefundPixUseCase(irefund, rep, QUEUES)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	c.Handler(http.DefaultServeMux)

	log.Printf("api running port %d\n", API_PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", API_PORT), nil); err != nil {
		log.Fatalln(err)
	}
}
