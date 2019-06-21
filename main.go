package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"first"` // want to change this to `json:"name"`
	tag  string `json:"-"`
	Another
}

type Another struct {
	Address string `json:"address"`
}

type MyUser struct {
	U User
}

func (u MyUser) MarshalJSON() ([]byte, error) {
	// encode the original
	m, _ := json.Marshal(u.U)

	// decode it back to get a map
	var a interface{}
	json.Unmarshal(m, &a)
	fmt.Println("-----")
	fmt.Println(a)
	fmt.Println("-----")
	b := a.(map[string]interface{})

	// Replace the map key
	b["name"] = b["first"]
	delete(b, "first")

	// Return encoding of the map
	return json.Marshal(b)
}

func main() {
	anoth := Another{"123 Jennings Street"}
	u := User{1, "Ken Jennings", "name", anoth}
	e := json.NewEncoder(os.Stdout)
	e.Encode(u)
	e.Encode(MyUser{u})

	var str_emp = []byte("{\"Name\":\"Rachel\",\"Age\":24,\"Salary\":344444}")

	type Response struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Salary int    `json:"salary"`
	}
	bytes := []byte(str_emp)
	var res Response
	json.Unmarshal(bytes, &res)

	

}

