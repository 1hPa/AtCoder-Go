package main
import (
  "fmt"
  "strconv"
  "math"
)

func main(){
  var a,b string
  fmt.Scan(&a, &b)
  var s int
  s, _ = strconv.Atoi(a+b)
  var n int
  n = int(math.Sqrt(float64(s)))
  if n*n == s{
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}
