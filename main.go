package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		fmt.Fprintf(w, "Hello, this is %s", host)
	})
	portStr := os.Getenv("PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080
	}
	listen := fmt.Sprintf(":%d", port)
	fmt.Printf("listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, nil))

}
