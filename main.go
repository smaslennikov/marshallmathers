package main

import (
    "fmt"
    "os"
    "log"
    "io"
    "crypto/rand"

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
    Salt string `yaml:"salt"`
    Locked bool `yaml:"locked"`
}

func main() {
    list := UserList{}
    salt := []byte("")

    // generate and set a salt if user specifies a flag
    // TODO: use cobra to check flags and inform user of their purposes
    if len(os.Args) > 1 {
        fmt.Println("Using a random salt for hashing passwords")
    }

    // read in supplied sample yaml
    file, err := ioutil.ReadFile("sample.yaml")
    if err != nil {
        log.Fatalf("file.read error: %v", err)
    }

    // unmarshal supplied sample yaml
    err = yaml.Unmarshal(file, &list)
    if err != nil {
        log.Fatalf("unmarshal error: %v", err)
        return
    }

    for i := range list.Users {
        // lock each user
        list.Users[i].Locked = false

        // generate and set a salt if user specifies a flag
        if len(os.Args) > 1 {
            // TODO: variably set length of salt
            salt = make([]byte, 128)
            _, err = io.ReadFull(rand.Reader, salt)
            if err != nil {
                log.Fatalf("random salt error: %v", err)
            }

            list.Users[i].Salt = string(salt)
        }

        list.Users[i].Pass = string(pbkdf2.Key([]byte(list.Users[i].Pass), salt, 4096, 32, sha1.New))
    }

    // TODO: obviously, if salt isn't required, it's still exported when marshalled. Fix this
    d, err := yaml.Marshal(&list)
    if err != nil {
        log.Fatalf("marshal error: %v", err)
        return
    }

    // TODO: variably set output filename via arguments
    outfile, err := os.Create("output.yaml")
    if err != nil {
        log.Fatalf("outfile creation error: %v", err)
        return
    }

    _, err = io.WriteString(outfile, string(d))
    if err != nil {
        log.Fatalf("outfile write error: %v", err)
        return
    }

    outfile.Sync()
}
