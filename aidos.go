/* Copyright (C)  2016 nokaa <nokaa@cock.li>
 * This software is licensed under the terms of the
 * GNU Affero General Public License. You should have
 * received a copy of this license with this software.
 * The license may also be found at https://gnu.org/licenses/agpl.txt
 */

package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"net/http"
	"strings"
)

const (
	PORT = ":5000"
)

var db *bolt.DB

func main() {
	var err error
	db, err = bolt.Open("aidos.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", SearchHandler)

	fmt.Println("Listening at localhost" + PORT)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		db.Close()
		return
	}
	db.Close()
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("q")
	if strings.HasPrefix(search, "!") {
		bang, search := split(search)

		site := checkDB(bang)
		if site == nil && bang == "" {
			http.Redirect(w, r, "https://duckduckgo.com/?q="+search, 302)
		} else if site == nil {
			http.Redirect(w, r, "https://duckduckgo.com/?q=!"+bang+" "+search, 302)
		} else {
			http.Redirect(w, r, string(site)+search, 302)
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

func checkDB(bang string) []byte {
	var valid []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bangs"))
		if bucket == nil {
			return fmt.Errorf("Bucket bangs not found!")
		}

		valid = bucket.Get([]byte(bang))

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return valid
}
