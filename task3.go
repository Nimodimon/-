package main

import (
  "fmt"
  "math/big"
)

func main()  {
  const num = 600851475143

  if num == 2 {
    fmt.Println(num)
  } else {
    biggest_div := 1

    for i := 3; i <= num; i+=2{
      if num % i == 0 && big.NewInt(2).ProbablyPrime(0) && i > biggest_div{
        biggest_div = i
      }
    }

    fmt.Println(biggest_div)
  }
}
