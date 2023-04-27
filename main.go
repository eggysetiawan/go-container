package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	fmt.Println("Server is running at localhost:8800")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			"You are visiting go container",
			http.StatusOK,
		}

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe("0.0.0.0:8800", r)

}
