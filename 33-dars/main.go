package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/read", ReadText)
	http.HandleFunc("/reverseStr/", ReverseString)
	http.HandleFunc("/golang/", ProgrammingLanguage)
	http.HandleFunc("/devops/", Developer)

	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		panic(err)
	}
}

func ReadText(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Salom Dunyo"))
	if err != nil {
		panic(err)
	}
}

func ReverseString(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Salom Dunyo"))
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(data); i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}

func ProgrammingLanguage(w http.ResponseWriter, r *http.Request) {
	languages := map[string]string{"100": "GO", "101": "Java", "102": "Python"}
	w.Write([]byte(languages[strings.TrimPrefix(r.URL.Path, "/golang/")]))
}

type developer struct {
	Id        int
	Name      string
	Direction string
	Salary    float64
}

func Developer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("XATOLIK"))
		panic(err)
	}
	dev := developer{}
	json.Unmarshal(body, &dev)
	fmt.Println(dev)
}
