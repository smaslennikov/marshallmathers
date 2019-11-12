package main

import (
    "fmt"
    "log"

    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type UserList struct {
    Users []User `yaml:"users"`
}

type User struct {
    Name string `yaml:"username"`
    Pass string `yaml:"password"`
    Locked bool `yaml:"locked"`
}

func main() {
    list := UserList{}

    file, err := ioutil.ReadFile("sample.yaml")
    if err != nil {
        log.Fatalf("file.read error: %v", err)
    }

    //fmt.Println(string(file))

    err = yaml.Unmarshal(file, &list)
    if err != nil {
        log.Fatalf("unmarshal error: %v", err)
        return
    }

    fmt.Printf("%v\n\n", list)
}
