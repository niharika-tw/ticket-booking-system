package app

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/book", bookHandler())
	http.HandleFunc("/", pingHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}
	fmt.Println("Server Started")
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}

func bookHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Ticket Created")
		writer.WriteHeader(http.StatusCreated)
	}
}
