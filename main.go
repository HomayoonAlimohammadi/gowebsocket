package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HomayoonAlimohammadi/gowebsocket/route"
)

func setupRoutes() {
	http.HandleFunc("/", route.Home)
	http.HandleFunc("/ws", route.WebSocket)
}

func main() {
	fmt.Println("Running server...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
