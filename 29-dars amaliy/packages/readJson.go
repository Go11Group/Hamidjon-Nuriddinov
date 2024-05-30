package packages

import (
	"encoding/json"
	users "mymod/packages/Users"
	"os"
)

func ReadJson() []users.User {
	users := []users.User{}
	file, err := os.OpenFile("data.json", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		panic(err)
	}
	return users
}