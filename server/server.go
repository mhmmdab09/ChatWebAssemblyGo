package main

import (
	"log"
	"net/http"
)

const (
	AddSrv       = ":8090"
	TemplatesDir = "../app/html/"
)

func main() {
	fileSrv := http.FileServer(http.Dir(TemplatesDir))

	err := http.ListenAndServe(AddSrv, fileSrv)

	if err != nil {
		log.Fatalln(err)
	}
}
