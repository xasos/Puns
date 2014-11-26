package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"math/rand"
	"encoding/json"
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

       rand.Seed(time.Now().UnixNano())
       fmt.Println(rand.Intn(100))
    
    m.Get("/api/random", func() string {
	return "test"
    })

    m.Run()
}
