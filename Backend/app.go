package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"encoding/json"
	"io/ioutil"
)

func main() {
    m := martini.Classic()

    file, e := ioutil.ReadFile("./puns.json")    
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    m.Get("/api/random", func() string {
        return "test"
    })

    m.Run()
}
