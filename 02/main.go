package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Data struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	datas := make([]*Data, 0, 100)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		res := datas

		resBody, err := json.Marshal(res)
		if err != nil {
			body := &Error{
				Code:    500,
				Message: "internal server error",
			}
			resBody, err := json.Marshal(body)
			if err != nil {
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
			w.WriteHeader(500)
			w.Write(resBody)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)

	})
	r.Post("/data", func(w http.ResponseWriter, r *http.Request) {
		var req Data

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			body := &Error{
				Code:    400,
				Message: "Bad request",
			}
			resBody, err := json.Marshal(body)
			if err != nil {
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
			w.WriteHeader(400)
			w.Write(resBody)
			return
		}

		datas = append(datas, &req)
		log.Println(datas)

		resBody, _ := json.Marshal(req)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	})

	http.ListenAndServe(":8080", r)
}
