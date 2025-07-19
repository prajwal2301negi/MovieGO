package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prajwal2301negi/db/router"
)

func main() {
	fmt.Println("Mongo BD in GO")
	r := router.Router()
	fmt.Println("Server is getting started ......")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ....")
}
