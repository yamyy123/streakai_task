package main

import (
	"fmt"
	"net/http"
	"streakai/handler"
)

func main() {
	http.HandleFunc("/find-paths", handler.DfsHandler)
	fmt.Println("Server Started running on 8080....")
	http.ListenAndServe(":8080", nil)
}
