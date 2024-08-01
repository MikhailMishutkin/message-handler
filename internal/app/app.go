package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/segmentio/kafka-go"
	"log"
	"message_handler/configs"
	kafkarepo2 "message_handler/internal/repository/kafka"
	"message_handler/internal/repository/postgres"
	httpservice "message_handler/internal/service/http"
	httphandler "message_handler/internal/transport/http"
	"net/http"
)

type Server struct {
	Router *gin.Engine
}

func StartService(conf configs.Config) error {
	s := &Server{
		Router: gin.Default(),
	}
	//postgres
	db, err := NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	//kafka
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//httpserver
	repo := postgres.NewRepo(db)
	kafkarepo := kafkarepo2.NewKafkaRepo(conn)
	messageService := httpservice.NewMessageService(repo, kafkarepo)
	handler := httphandler.NewHTTPMessageHandle(messageService)
	handler.RegisterMessage(s.Router)

	log.Println("Starting MessageService at port: 8080")
	return http.ListenAndServe(":8080", s)
}

func NewDB() (*pgx.Conn, error) {
	c, err := configs.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.Conn)

	db, err := pgx.Connect(context.Background(), psqlInfo)

	m, err := migrate.New(
		"file://../message_handler/migrations",
		"postgres://"+c.Migrate,
		//root:root@localhost:5444/time_tracker?sslmode=disable",
	)
	if err != nil {
		log.Println(err)
		return db, fmt.Errorf("can't automigrate: %v\n", err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
		fmt.Errorf("%v\n", err)
	}
	return db, err
}

// ServeHTTP
func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}
