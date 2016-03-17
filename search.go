package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SearchHandler)

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, `<html><body>This server is POST only. Please
		<a href="https://git.nokaa.moe/nokaa/search">see the repository</a>
		for more information.</body></html>`)
		return
	}
}
