package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "code=500")
	} else {
		query := request.FormValue("name")
		if query == "student1" {
			fmt.Fprintf(writer, "90")
			writer.WriteHeader(200)
		} else if query == "student2" {
			fmt.Fprintf(writer, "80")
			writer.WriteHeader(200)
		} else if query == "" {
			writer.WriteHeader(500)
		} else {
			writer.WriteHeader(404)
		}
	}
}

func main() {
	http.HandleFunc("/mark", viewHandler)
	err := http.ListenAndServe("localhost:8082", nil)
	log.Fatal(err)
}
