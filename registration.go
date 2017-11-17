package main

import (
  "io/ioutil"
)

import (
  "encoding/json"
  "fmt"
  "os"
)


func registration(){

  var user User

  fmt.Println("Registration")
  fmt.Println("")
  fmt.Print("Please tell me your name : ")
  fmt.Scan(&user.Name)
  fmt.Print("Your Email : ")
  fmt.Scan(&user.Email)
  fmt.Print("Your Phone : ")
  fmt.Scan(&user.Phone)
  fmt.Print("Your Password : ")
  fmt.Scan(&user.Password)

  fmt.Println(user)

  user.Save("./user_registration.json")


}


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
