package main
import(
  "fmt"
  "math"
)

func main(){
  var r, x, y float64

  fmt.Scanln(&r, &x, &y)
  distance := math.Sqrt(x*x + y*y)

  if distance < r{
    fmt.Println("2")
  }else{
    fmt.Println(math.Ceil(distance / r))
  }
}
