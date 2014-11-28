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

	// Import Puns.json file and serialize into go struct
	file, err := ioutil.ReadFile("./puns.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	type Puns []struct {
		Pun string `json:"Pun"`
	}	

	var p Puns
	err = json.Unmarshal([]byte(file), &p)
	if err != nil {
		fmt.Println("error:", err)
	}
		
	m.Get("/api", func() string {
		return "{\"message\":\"Welcome! Please refer to the docs for API calls\"}"
	})

	m.Get("/api/random", func() string {
		rand.Seed(time.Now().UnixNano())
	    randomNum := rand.Intn(100)

	    b, err := json.Marshal(p[randomNum])
	    if err != nil {
	    	fmt.Println("error:", err)
	    }
	    
	    // convert byte array into string
	    s := string(b[:])
	    fmt.Printf(s)
		return s
	})

	m.Run()
}
