package main

import (
	"http-todo-list/server"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/add", server.Add)
	http.HandleFunc("/view", server.View)
	http.HandleFunc("/edit", server.Edit)

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatal("ListenAndServer error:", err)
	}
}

