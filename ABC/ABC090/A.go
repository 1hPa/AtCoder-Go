package main
import "fmt"

func main(){
  var a,b,c string
  fmt.Scan(&a, &b, &c)
  fmt.Printf("%s%s%s", a[0:1], b[1:2], c[2:3])
}
