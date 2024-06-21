package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// URL ni olish
	url := r.URL.String()
	fmt.Println("Received URL:", url)
	// URL ni qaytarish (optional)
	fmt.Fprintf(w, "URL: %s", url)

	URL := "http://localhost:8080" + url

	resp, err := http.Get(URL)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Fprintf(w,string(body))
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server started at :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))

}