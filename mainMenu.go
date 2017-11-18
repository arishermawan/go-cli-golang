package main

import "fmt"

func mainMenu(Opts map[string]string) (i int) {

  fmt.Println("Welcome to Go-CLI!")
  fmt.Println("")

  fmt.Println("Main Menu")
  fmt.Println("1. View Profile")
  fmt.Println("2. Order Go-Ride")
  fmt.Println("3. View Order History")
  fmt.Println("4. Exit")

  fmt.Println("Enter your option: ")
  fmt.Scan(&i)

  return

}
