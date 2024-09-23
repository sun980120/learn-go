package main

import (
	"fmt"
	"net/http"

	router "gin-api/router"
)

func main() {
	printBanner()
	router.StartRouter()
	// http.HandleFunc("/", index)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
}

func printBanner() {
	fmt.Println("=======================================================")
	fmt.Println("=================== START SERVER ======================")
	fmt.Println("=======================================================")
}
