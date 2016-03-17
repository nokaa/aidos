package main

import (
	"fmt"
	"net/http"
	"strings"
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

	search := r.FormValue("q")
	if strings.HasPrefix(search, "!") {
		site, search := split(search)
	} else {
		http.Redirect(w, r, "https://duckduckgo.com/?q="+search, 302)
	}
}

func split(search string) (string, string) {
	j := 1
	var site string
	var searchString string

	for i := j; i < len(search); i++ {
		if search[i] == ' ' {
			j = i + 1
			break
		} else {
			site += string(search[i])
		}
	}

	for i := j; i < len(search); i++ {
		searchString += string(search[i])
	}

	return site, searchString
}
