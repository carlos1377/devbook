package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carlos1377/devbook/api/config"
	"github.com/carlos1377/devbook/api/router"
)

func main() {
	config.Load()
	fmt.Printf("Running on %d! 🚀", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
