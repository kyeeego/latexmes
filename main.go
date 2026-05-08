package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kyeeego/latexmes/config"
	"github.com/kyeeego/latexmes/delivery/http"
	"github.com/kyeeego/latexmes/delivery/ws"
	"github.com/kyeeego/latexmes/repository"
	"github.com/kyeeego/latexmes/server"
	"github.com/kyeeego/latexmes/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conf := config.Init()

	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.DBUsername, conf.DBPassword),
		Host:     fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()))
	if err != nil {
		log.Fatalf("Could not connect to database: %e\n", err)
	}

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	s3client := s3.NewFromConfig(cfg)

	repos := repository.New(db)
	services := service.New(repos, s3client)
	wsHandler := ws.NewHandler(services)
	handler := http.New(services, wsHandler)

	s := &server.Server{}

	if err := s.Run(conf.Port, handler.Init()); err != nil {
		log.Fatalf("Got an error while trying to start the server: %e\n", err)
	} else {
		log.Printf("Server is active on port :%d\n", conf.Port)
	}
}
