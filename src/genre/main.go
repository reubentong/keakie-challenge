package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/username/reubentong/genre"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbSource = "genres-db.json"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "genre",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	content, err := ioutil.ReadFile(dbSource)
	if err != nil {
		logger.Log("Error when opening file: ", err)
	}

	var genres []genre.Genre

	err = json.Unmarshal(content, &genres)
	if err != nil {
		logger.Log("Error during unmarshalling of db: ", err)
	}

	flag.Parse()
	ctx := context.Background()
	var srv genre.Service
	{
		repository := genre.NewRepository(genres, logger)

		srv = genre.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := genre.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := genre.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
