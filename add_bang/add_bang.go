package main

import (
	"bufio"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

func main() {
	db, err := bolt.Open("aidos.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Please enter a bang (i.e. `g` not `!g`): ")
	consoleReader := bufio.NewReader(os.Stdin)
	bang, err := consoleReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	bang = bang[:len(bang)-1] // Trim newline character
	fmt.Printf("Please enter the URL for this bang: ")
	url, err := consoleReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	url = url[:len(url)-1] // Trim newline character

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("bangs"))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(bang), []byte(url))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
