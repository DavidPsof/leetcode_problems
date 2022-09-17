package server

import (
	"fmt"
	"github.com/DavidPsof/leetcode_problems/backend/config"
	"github.com/DavidPsof/leetcode_problems/backend/server/http"
	"github.com/DavidPsof/leetcode_problems/backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Server struct {
	cfg     config.Conf
	app     *fiber.App
	service *services.Service
}

func NewServer(conf config.Conf) (*Server, error) {
	app := fiber.New()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.PostgresConfig.Host, conf.PostgresConfig.Port, conf.PostgresConfig.User, conf.PostgresConfig.Pass, conf.PostgresConfig.DBName, conf.PostgresConfig.ModeSSL))
	if err != nil {
		log.Fatalln(err)
	}

	srvs := services.NewService(db)

	return &Server{
		app:     app,
		cfg:     conf,
		service: srvs,
	}, nil
}

func (s *Server) Run() {
	http.InitRoutes(s.app)

	log.Fatal(s.app.Listen(fmt.Sprintf(":%v", s.cfg.HttpServerConfig.Port)))
}
