package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	r := router.Generate()

	utils.LoadTemplates()

	fmt.Println("Application Started. Listenning Port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
