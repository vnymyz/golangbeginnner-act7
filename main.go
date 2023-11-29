package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	PORT := 9068 // Ganti 2 digit terakhir npm

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Starting server on port %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}