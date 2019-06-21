package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
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
	b := a.(map[string]interface{})
	fmt.Println(b)
	// Replace the map key
	b["identificacion"] = b["id"]
	b["nombre"] = b["name"]
	b["direccion"] = b["address"]

	delete(b, "id")
	delete(b, "name")
	delete(b, "address")

	// Return encoding of the map
	
	//fmt.Println(json.Unmarshal(bytes, &b))

	return json.Marshal(b)
}


func main() {
	u := User{11212, "Ken Jennings", "quito"}
	data := []byte(`
            {
                "id": 14,
                "name": "ruben,
                "address": "ruben"
            }
        `)
	var user User
	json.Unmarshal(data, &user)
	fmt.Println("*************")
	fmt.Println(user)
	fmt.Println("*************")

	/*
	if err != nil {
                fmt.Println("error")
                os.Exit(1)
        }*/
	
	e := json.NewEncoder(os.Stdout)
	e.Encode(u)
	e.Encode(MyUser{u})
}

