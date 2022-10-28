package main

import (
	"fmt"
)

func main() {
	var opt int
  
  fmt.Println("Figure Option : \n 1) Circle \n 2) Square \n 3) Triangle \n 4) Rectangle \n")
  fmt.Println("Insert your choice below : ")
  fmt.Scan(&opt)

  var area float32
  var circumference float32

  if(opt == 1){
    var r float32
    var pi float32
    pi = 3.14

    fmt.Print("\n=== CIRCLE ===\n radius : ")
    fmt.Scan(&r)

    area = pi*r*r
    circumference = 2*pi*r
  } else {
    fmt.Println("\nlooks like you choose the wrong option :((\n ")  
  }

  fmt.Println("\n=== Result ===\n area : ", area, "\n circumference : ", circumference)
  
}
