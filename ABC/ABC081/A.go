package main
import (
  "fmt"
)

func main(){
  var count int
  var str string
  fmt.Scan(&str)
  if str[0] == '1'{count++}
  if str[1] == '1'{count++}
  if str[2] == '1'{count++}
  fmt.Println(count)
}
