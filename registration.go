package main

import (
  "fmt"
)


func registration(Opts map[string]string){

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

  Opts["name"] = user.Name
  Opts["email"] = user.Email
  Opts["phone"] = user.Phone
  Opts["password"] = user.Password
  
  fmt.Println(user)

  user.Save("./user_registration.json")
}
