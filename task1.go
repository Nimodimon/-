package main

import "fmt"

func main()  {
  sum := 0

  for i := 1; i < 1000; i += 1 {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }

  fmt.Print(sum)
}
