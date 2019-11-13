package main

import (
    "fmt"
    "log"

    "golang.org/x/crypto/pbkdf2"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "crypto/sha1"
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

    err = yaml.Unmarshal(file, &list)
    if err != nil {
        log.Fatalf("unmarshal error: %v", err)
        return
    }

    for i := range list.Users {
        list.Users[i].Locked = false
        salt := ""
        //fmt.Println(list.Users[i].Pass)
        list.Users[i].Pass = string(pbkdf2.Key([]byte(list.Users[i].Pass), []byte(salt), 4096, 32, sha1.New))
    }

    fmt.Printf("%v\n\n", list)
}
