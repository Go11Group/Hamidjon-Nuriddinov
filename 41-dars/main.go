package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/users/readUser/68ec648a-cb75-4fdc-bdea-6680fcd8eaab")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	resp, err = http.Get("http://localhost:8080/users/getAllUsers")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	type User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Birthday string `json:"birthday"`
	}
	user := User{Name: "Anvar", Email: "anvar@gmail.com", Password: "12354", Birthday: "2024-06-05"}
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	// resp, err = http.Post("http://localhost:8080/users/createUser", "JSON", bytes.NewBuffer(data))
	// if err != nil {
	// 	panic(err)
	// }
	// body, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))

// 	client := http.Client{}
// 	req, err := http.NewRequest("PUT", "http://localhost:8080/users/updateUser/eaf8f40b-a8ca-4764-aff7-90ac763333db", bytes.NewBuffer(data))
// 	if err != nil{
// 		panic(err)
// 	}
// 	resp, err = client.Do(req)
// 	if err != nil{
// 		panic(err)
// 	}
// 	body, err = io.ReadAll(resp.Body)
// 	if err != nil{
// 		panic(err)
// 	}
// 	fmt.Println(string(body))
// 

client := http.Client{}
req, err := http.NewRequest("DELETE", "http://localhost:8080/users/deleteUser/eaf8f40b-a8ca-4764-aff7-90ac763333db", bytes.NewBuffer(data))
if err != nil{
	panic(err)
}
resp, err = client.Do(req)
if err != nil{
	panic(err)
}
body, err = io.ReadAll(resp.Body)
if err != nil{
	panic(err)
}
fmt.Println(string(body))

}


