package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("hello.html"))

	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
