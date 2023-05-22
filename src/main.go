package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	mux.HandleFunc("/student", StudentHandler)

	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"민혁", 28, 100}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":5000", MakeWebHandler())
}
