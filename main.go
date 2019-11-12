package main

import (
    "fmt"
    "log"

    "gopkg.in/yaml.v2"
)

var data = `username: test
password: badpassword
locked: false`

type User struct {
    Name string `yaml:"username"`
    Pass string `yaml:"password"`
    Locked bool `yaml:"locked"`
}

func main() {
    //fmt.Println(yaml)

    u1 := User{}

    err := yaml.Unmarshal([]byte(data), &u1)
    if err != nil {
        log.Fatalf("error: %v", err)
        return
    }

    fmt.Printf("--- u1:\n%v\n\n", u1)
}
