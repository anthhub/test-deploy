package main

import (
	"fmt"
	"net/http"
	"os"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	myhostname, _ := os.Hostname()
	fmt.Fprintln(w, "go-server--Hostname:", myhostname)
}

func main() {
	const port string = "8000"
	fmt.Println("Server listening on port", port)
	http.HandleFunc("/", hostnameHandler)
	http.ListenAndServe(":"+port, nil)
}
