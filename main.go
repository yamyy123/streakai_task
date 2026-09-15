package main

import (
	"fmt"
	"net/http"
	"streakai/handler"
)

func main() {
	http.HandleFunc("/find-paths", handler.DfsHandler)
	fmt.Println("Server Started running on 8080....")
	fmt.Println("To find paths, use the endpoint: http://localhost:8080/find-paths")
	http.ListenAndServe(":8080", nil)
}
