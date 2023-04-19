package main

import (
	"context"
	"fmt"
	"github.com/mereiamangeldin/transaction/config"
	"github.com/mereiamangeldin/transaction/repostiory"
	"github.com/mereiamangeldin/transaction/service"
	"github.com/mereiamangeldin/transaction/transport/http"
	"github.com/mereiamangeldin/transaction/transport/http/handler"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down:%s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := repostiory.New(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		log.Fatal(err.Error())
	}
	h := handler.NewManager(conf, svc)
	srv := http.NewServer(conf, h)
	return srv.Run(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
