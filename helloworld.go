package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := "localhost:9000"
	fmt.Printf("Server started at %s\n", address)
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
}
