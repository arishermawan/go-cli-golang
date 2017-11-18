package main

import "fmt"

func main() {

  user_data := UserFromFile("./user_registration.json")

  locationx := loadLocation()

  dist := Distance(locationx[0].Coord, locationx[2].Coord)

  fmt.Println(dist)

  fmt.Println("")

  for _, e := range locationx {
    // fmt.Print(i)
    // fmt.Print(" ")
    fmt.Print(e.Coord)
  }

  fmt.Println("")



  var Opts = map[string]string {
    "name": user_data.Name,
    "email": user_data.Email,
    "phone": user_data.Phone,
    "password": user_data.Password,
  }


  if Opts["name"] == "" {
    registrationController(Opts)
  }else{
    loginController(Opts)
  }

    mainMenuController(Opts)
}


func registrationController(opts map[string]string) {
  registration(opts)
}

func loginController(opts map[string]string){
  halt := false
  for halt == false {
    login(opts)

    if opts["email"]==opts["email_login"] && opts["password"]==opts["password_login"]  {
      halt = true
      fmt.Println("Anda Login")
    }
    fmt.Println("Anda Salah")
  }


}

func viewProfileController(opts map[string]string){
  input := viewProfile(opts)
  if input == 1 {
    registrationController(opts)
    fmt.Println("Your Profile Updated")
    viewProfileController(opts)
  }else{
    mainMenuController(opts)
  }
}

func orderGoRideController(opts map[string]string){
  fmt.Println("orderGoRideController")
  orderGoRide(opts)
  coord1 :=searchLocation(opts["pickup"])
  coord2 :=searchLocation(opts["destination"])
  fmt.Println(coord1)
  fmt.Println(coord2)
  cost := Distance(coord1, coord2) * 1500.0
  fmt.Println(cost)


}

func searchLocation(x string)(coord [2]int){
  location := loadLocation()
  for _, e := range location {
    if x == e.Name {
      coord = e.Coord
    }
  }
  return
}

func viewOrderHistoryController(opts map[string]string){
  fmt.Println("orderGoRideController")
}


func mainMenuController(opts map[string]string){

  var input = mainMenu(opts)

  switch input{
  case 1:
    viewProfileController(opts)

  case 2:
    orderGoRideController(opts)

  case 3:
    viewOrderHistoryController(opts)
  }
}


func loadLocation()(location []Location) {
  location = LocationFromFile("./locations.json")
  return
}

func loadOrders()(orders []Orders) {
  orders = OrderFromFile("./orders.json")
  return
}
