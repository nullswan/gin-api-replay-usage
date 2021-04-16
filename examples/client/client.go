package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Just to let server start if we use .sh test script
	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 10; i++ {
		go func() {
			req, err := http.NewRequest("GET", "http://127.0.0.1:3000/", bytes.NewReader(nil))
			if err != nil {
				fmt.Printf("failed to create request: %s\n", err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("failed to to call: %s\n", err)
			}
			defer res.Body.Close()

		}()
		time.Sleep(100 * time.Millisecond)
	}
}
