package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	data := Response{Message: "Welcome to the Go REST API!"}

	var buf bytes.Buffer
	j := json.NewEncoder(&buf)

	if err := j.Encode(&data); err != nil {
		log.Fatal(err)
	}

	res := buf.String()

	log.Println(res)
	_, err := fmt.Fprint(w, res)

	if err != nil {
		log.Fatal(err)
	}
}

func SetupServer() *mux.Router {
	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

func main() {
	r := SetupServer()
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
}
