package main

import "fmt"


func orderGoRide(opts map[string]string) {

  var pickup, destination string

  fmt.Print("Your Pick Up Location : ")
  fmt.Scan(&pickup)
  fmt.Print("Your Destination : ")
  fmt.Scan(&destination)

  opts["pickup"] = pickup
  opts["destination"] = destination

}
