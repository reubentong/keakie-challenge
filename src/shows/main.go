package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/reubentong/shows/shows"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbSource = "shows-db.json"

func main() {
	var httpAddr = flag.String("http", ":8081", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "shows",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	content, err := ioutil.ReadFile(dbSource)
	if err != nil {
		logger.Log("Error when opening file: ", err)
	}

	var showSlice []shows.Show

	err = json.Unmarshal(content, &showSlice)
	if err != nil {
		logger.Log("Error during unmarshalling of db: ", err)
	}

	flag.Parse()
	ctx := context.Background()
	var srv shows.Service
	{
		repository := shows.NewRepository(showSlice, logger)

		srv = shows.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := shows.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := shows.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
