package main

import (
  "io/ioutil"
)

import (
  "encoding/json"
  "fmt"
  "os"
)

func login() {

 var email, password string

  fmt.Println("Login")
  fmt.Println("")
  fmt.Print("Your Email : ")
  fmt.Scan(&email)
  fmt.Print("Your Passsword : ")
  fmt.Scan(&password)

  userFile := UserFromFile("./user_registration.json")

  if email == userFile.Email && password == userFile.Password{
    fmt.Println("Anda Login")
  }else{
    fmt.Println("Anda Salah")
  }


}


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
