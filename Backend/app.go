package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	m := martini.Classic()

	type Puns struct {
		Pun string
	}

	// Import Puns.json file and serialize into go struct
	file, e := ioutil.ReadFile("./puns.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}

	p := Puns{}
	json.Unmarshal([]byte(file), &p)

	m.Get("/api/random", func() string {
		rand.Seed(time.Now().UnixNano())
	        fmt.Println(rand.Intn(100))
		return "test"
	})

	m.Run()
}
