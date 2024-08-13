package main

import "net/http"

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
}
