package main
import (
  "bufio"
  "fmt"
  "strings"
  "os"
)

func main(){
  var lowerString string
  var trimmedString string
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Please provide the string that you would like to test for the presence of 'i','a','n' : ")
  input, _ := consoleReader.ReadString('\n')
  lowerString = strings.ToLower(input)
  trimmedString = strings.TrimSpace(lowerString)

  /*fmt.Println(input)
  fmt.Println(fmt.Sprintf("%T",input))
  fmt.Println(fmt.Sprintf("%T",lowerString))
  fmt.Println(lowerString)
  fmt.Println(trimmedString)*/

  if strings.HasPrefix(trimmedString,"i") && strings.Contains(trimmedString,"a") && strings.HasSuffix(trimmedString,"n"){
    fmt.Println("Found!")
  }else{
    fmt.Println("Not Found!")
}
}
