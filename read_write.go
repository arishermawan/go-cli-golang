package main


import (
  "io/ioutil"
)

import (
  "encoding/json"
  "fmt"
  "os"
)


//save to json
func (u User) Save(f string) error {
  return ioutil.WriteFile(f, u.toJson(), 0644)
}

// Convert to json
func (u User) toJson() []byte {
  bs, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

// Read from file
func UserFromFile(f string) User {
  bs, err := ioutil.ReadFile(f)
  if err != nil {
    return User{}
  }
  return userFromJson(bs)
}

// Read from json
func userFromJson(bs []byte) User {
  var u User
  err := json.Unmarshal(bs, &u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return u
}
