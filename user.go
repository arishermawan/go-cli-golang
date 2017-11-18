package main

import "math"

type User struct{
  Name string
  Email string
  Phone string
  Password string
}

type Location struct{
  Name string
  Coord [2]int
}

type Order struct{
  Timestamp string
  Origin string
  destination string
  est_price float64
}

func Distance(x, y [2]int)(dist float64){
  dist = math.Sqrt(float64( (y[0] - x[0])*(y[0] - x[0]) + (y[1] - x[1])*(y[1] - x[1]) ))
  return
}
