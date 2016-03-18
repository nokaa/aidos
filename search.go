package main

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	PORT = ":5000"
)

func main() {
	http.HandleFunc("/", SearchHandler)

	fmt.Println("Listening at localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, `<html><body>This server is POST only. Please
		<a href="https://git.nokaa.moe/nokaa/search">see the repository</a>
		for more information.</body></html>`)
		return
	}

	search := r.FormValue("q")
	if strings.HasPrefix(search, "!") {
		site, search := split(search)

		switch site {
		case "w":
			http.Redirect(w, r, "https://en.wikipedia.org/w/index.php?search="+search, 302)
		case "wt":
			http.Redirect(w, r, "https://en.wiktionary.org/w/index.php?search="+search, 302)
		case "gh":
			http.Redirect(w, r, "https://github.com/search?q="+search, 302)
		case "g":
			http.Redirect(w, r, "https://encrypted.google.com/search?q="+search, 302)
		}
		if site == "" {
			http.Redirect(w, r, "https://duckduckgo.com/?q="+search, 302)
		} else {
			http.Redirect(w, r, "https://duckduckgo.com/?q=!"+site+" "+search, 302)
		}
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
