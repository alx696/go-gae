package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var ip string

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(fmt.Sprintf("<h1 style=\"text-align: center;\">%s</h1>", ip)))
	})

	http.HandleFunc("/ip", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			_, _ = writer.Write([]byte(ip))
		case "POST":
			ip = request.FormValue("ip")
			writer.WriteHeader(http.StatusOK)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "60000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
