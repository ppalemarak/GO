package main
import "fmt"


func main(){
  var x float64
  fmt.Printf("Please provide the floating point number that you would like to convert to an Integer :")
  fmt.Scan(&x)
  fmt.Println("The truncated number is" ,int(x))

}
