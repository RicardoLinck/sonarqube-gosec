package main

import (
	"crypto/md5"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health/{user}", healthHandler)

	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write(weakHash("Healthy"))
}

func weakHash(value string) []byte {
	hash := md5.New()
	return hash.Sum([]byte(value))
}
