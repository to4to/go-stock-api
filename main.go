package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/to4to/go-stock-api/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env File Load Failed")
	}

	r := router.Router()
	PORT := os.Getenv("PORT")

	fmt.Println("Starting Server")

	log.Fatal(http.ListenAndServe(PORT, r))
	fmt.Println("Server Successfully Started :)")
}
