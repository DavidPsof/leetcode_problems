package server

import (
	"context"
	"fmt"
	"github.com/DavidPsof/leetcode_problems/backend/config"
	"github.com/DavidPsof/leetcode_problems/backend/migrations"
	"github.com/DavidPsof/leetcode_problems/backend/server/http"
	"github.com/DavidPsof/leetcode_problems/backend/services"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type Server struct {
	cfg     config.Conf
	app     *fiber.App
	service *services.Service
}

func NewServer(conf config.Conf) (*Server, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancel()

	app := fiber.New()

	db, err := NewGORM(conf)
	if err != nil {
		return nil, err
	}

	m, err := NewMinIO(conf.MinioConfig)
	if err != nil {
		return nil, err
	}

	srvs := services.NewService(ctx, db, m)

	return &Server{
		app:     app,
		cfg:     conf,
		service: srvs,
	}, nil
}

func NewGORM(conf config.Conf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.PostgresConfig.Host, conf.PostgresConfig.Port, conf.PostgresConfig.User, conf.PostgresConfig.Pass, conf.PostgresConfig.DBName, conf.PostgresConfig.ModeSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	if err = migrations.Init(db); err != nil {
		return nil, err
	}

	return db, nil
}

func NewMinIO(conf config.MinioConfig) (*minio.Client, error) {
	mc, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return mc, nil
}

func (s *Server) Run() {
	h := http.NewHandler(s.service)

	h.InitRoutes(s.app)

	// TODO: поменять лог на нормальный

	log.Fatal(s.app.Listen(fmt.Sprintf(":%v", s.cfg.HttpServerConfig.Port)))
}
