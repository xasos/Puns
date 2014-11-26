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
    
    type jsonobject struct {
    Object ObjectType
    }

    file, e := ioutil.ReadFile("./puns.json")    
    if e != nil {
        fmt.Printf("File error: %v\n", e)
    }
    fmt.Printf("%s\n", string(file))
    
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)

    s1 := rand.NewSource(42)
    r1 := rand.New(s1)
    fmt.Print(r1.Intn(100))

    m.Get("/api/random", func() string {
        return "test"
    })

    m.Run()
}
