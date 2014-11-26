package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"math/rand"
	"encoding/json"
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
        s1 := rand.NewSource(42)
        r1 := rand.New(s1)
        fmt.Print(r1.Intn(99))

	return "test"
    })

    m.Run()
}
