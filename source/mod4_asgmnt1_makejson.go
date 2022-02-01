package main
import (
  "fmt"
  "encoding/json"
  "bufio"
  "os"
  "strings"
)

func main(){

    consoleReader := bufio.NewReader(os.Stdin)
    fmt.Print("What is your name? ")
    input, _ := consoleReader.ReadString('\n')
    var name = strings.TrimSpace(input)

    consoleReader1 := bufio.NewReader(os.Stdin)
    fmt.Print("What is your address? ")
    input1, _ := consoleReader1.ReadString('\n')
    var address = strings.TrimSpace(input1)

    var person = map[string]string{
      "name": name,
      "address": address,
    }
    //fmt.Println(person)

    barr,err := json.Marshal(person)
    if err == nil{
          fmt.Println(string(barr))
    }else {
      fmt.Println("Looks like sometthing went wrong .please check your inputs and try again. ")
    }

}
