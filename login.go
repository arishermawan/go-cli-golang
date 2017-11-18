package main

import (
  "fmt"
)

func login(opts map[string]string) {

  var email, password string

  fmt.Println("Login")
  fmt.Println("")
  fmt.Print("Your Email : ")
  fmt.Scan(&email)
  fmt.Print("Your Passsword : ")
  fmt.Scan(&password)

  opts["email_login"] = email
  opts["password_login"] = password
}
