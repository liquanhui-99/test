package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthZ", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success!"))
	})

	http.HandleFunc("/getHeader", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path: ", r.RequestURI)
		fmt.Println("ip: ", r.RemoteAddr)
		fmt.Println("statusCode: ", http.StatusOK)
		xApiKey := r.Header.Get("x-api-key")
		version := os.Getenv("VERSION")
		w.Header().Add("x-api-key", xApiKey)
		w.Header().Add("version", version)
		_, _ = w.Write([]byte("success"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("Started serve")
}
