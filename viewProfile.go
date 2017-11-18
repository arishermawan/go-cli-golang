package main

import (
  "fmt"
)


func viewProfile(opts map[string]string)(i int){

  fmt.Println("Your Profile")
  fmt.Println("")
  fmt.Print("Please tell me your name : ")
  fmt.Println(opts["name"])
  fmt.Print("Your Email : ")
  fmt.Println(opts["email"])
  fmt.Print("Your Phone : ")
  fmt.Println(opts["phone"])
  fmt.Print("Your Password : ")
  fmt.Println(opts["password"])

  fmt.Println("Please Choose a menu : ")
  fmt.Println("1. Edit Profile")
  fmt.Println("2. Back")

  fmt.Scan(&i)
  return

}
