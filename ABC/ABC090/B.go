package main
import "fmt"

func main(){
  var a,b,ans int
  fmt.Scan(&a, &b)
  for i := a; i <= b; i++ {
    if i/10000 == i%10 && i/1000%10 == i/10%10 {
      ans++
    }
  }
  fmt.Println(ans)
}
