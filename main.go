package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/castus/lights-api/api"
)

func main() {
	logger, _ := zap.NewProduction()
	log := logger.Sugar()
	defer logger.Sync()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		queryParameters := r.URL.Query()
		timestamp := queryParameters.Get("timestamp")
		parsedTime, err := strconv.ParseInt(timestamp, 10, 64)
		fmt.Println(parsedTime)
		if err != nil {
			panic(err)
		}
		unixTime := time.Unix(parsedTime, 0)
		lightValue := api.Get(unixTime)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Header().Set("Content-Type", "plain/text")
		_, _ = fmt.Fprintf(w, lightValue)
	})

	port := "8080"
	log.Infow("API server is running", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
