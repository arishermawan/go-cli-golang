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

// Read from file
func LocationFromFile(f string) []Location {
  bs, err := ioutil.ReadFile(f)
  if err != nil {
    return []Location{}
  }
  return LocationFromJson(bs)
}

// Read from json
func LocationFromJson(bs []byte) []Location {
  var u []Location
  err := json.Unmarshal(bs, &u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return u
}


// save json
func (u []Order) Save(f string) error {
  return ioutil.WriteFile(f, u.toJson(), 0644)
}

// Convert to json
func (u []Order) toJson() []byte {
  bs, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

func OrderFromFile(f string) []Order {
  bs, err := ioutil.ReadFile(f)
  if err != nil {
    return []Order{}
  }
  return OrderFromJson(bs)
}

// Read from json
func OrderFromJson(bs []byte) []Order {
  var u []Order
  err := json.Unmarshal(bs, &u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return u
}


// Read from file
// func LocationFromFile(f string) map[string] {
//   bs, err := ioutil.ReadFile(f)
//   if err != nil {
//     return map[string]{}
//   }
//   return LocationFromJson(bs)
// }

// // Read from json
// func LocationFromJson(bs []byte) map[string]{
//   var u map[string]
//   err := json.Unmarshal(bs, &u)
//   if err != nil {
//     fmt.Println(err)
//     os.Exit(1)
//   }
//   return u
// }
