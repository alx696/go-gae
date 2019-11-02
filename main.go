package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var ip string

func main() {
	//获取程序目录
	appDir, e := filepath.Abs(filepath.Dir(os.Args[0]))
	if e != nil {
		log.Fatal(e)
	}
	log.Println("程序目录:", appDir)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, e := template.ParseFiles(fmt.Sprint(appDir, "/template/index.html"))
		if e != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(e.Error()))
			return
		}

		_ = t.Execute(writer, ip)
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
