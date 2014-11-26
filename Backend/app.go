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

    file, e := ioutil.ReadFile("./puns.json")    
    if e != nil {
        fmt.Printf("File error: %v\n", e)
    }
    fmt.Printf("%s\n", string(file))
    
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)

    m.Get("/api/random", func() string {
        return "test"
    })

    m.Run()
}
