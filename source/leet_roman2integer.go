package main
import (
  "fmt"
  //"strconv"
  "bufio"
  "os"
  "strings"
  //"log"
)

/*func Swap(swap_slice []int, position int){
  x := swap_slice[position]
  swap_slice[position] = swap_slice[position + 1]
  swap_slice[position + 1 ] = x
}*/

func romanToInt(s string) int{
  /*roman_chars := []string{"I","V,"X","L","C","D","M"}
  for _, v := range s {
		if v == str {
			return true
		}*/

  size_string := len (s)
    current_value := 0
    skip_position := -1
  for i,_ := range s {
		//fmt.Printf("%c\n",s[i])
    var substract int
    //fmt.Println(fmt.Sprintf("%T", string(s[i])))
    var next_char string
    if skip_position == i{
      continue
    }
    if i < (size_string - 1){
    next_char = string(s[i+1])
  }
    //fmt.Println("The next char is:", next_char)
    switch string(s[i]){
    case "I":
        if next_char == "V" || next_char == "X" {
          skip_position = i+1
          if next_char == "V"{
          substract = 5
          } else if next_char == "X"{
            substract = 10
          }
          current_value = current_value + (substract - 1)
        }else{
          current_value = current_value + 1
        }
        //fmt.Print("I",i,current_value)

        case "V":
            current_value = current_value + 5
            //fmt.Print("V",i,current_value)

        case "X":
          if  next_char == "L" || next_char == "C" {
            skip_position = i+1
            if  next_char == "L"{
            substract = 50
            } else if next_char == "C"{
              substract = 100
            }
            current_value = current_value + (substract - 10)
          }else{
            current_value = current_value + 10
          }
          //fmt.Print("X",i,current_value)

        case "L":
          current_value = current_value + 50
          //fmt.Print("L",i,current_value)

        case "C":
          if  next_char == "D" || next_char== "M" {
            skip_position = i+1
            if next_char == "D"{
            substract = 500
            } else if next_char == "M"{
              substract = 1000
            }
            current_value = current_value + (substract - 100)
          }else{
            current_value = current_value + 100
          }
          //fmt.Print("C",i,current_value)
        case "D":
          current_value = current_value + 500
          ////fmt.Print("D",i,current_value)
        case "M":
          current_value = current_value + 1000
          //fmt.Print("M",i,current_value)

      }

      }
        return current_value
    }

func main(){

  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome ! This program converts Roman Numerals to integers.")
  fmt.Print("Please provide the Roman numeral you would like to convert to an integer. e.g.LVIII : ")
  fmt.Print("\n")
  input, _ := consoleReader.ReadString('\n')

  upperString := strings.ToUpper(input)
  trimmedString := strings.TrimSpace(upperString)
  fmt.Println(trimmedString)

  int_format := romanToInt(trimmedString)
  fmt.Println("The integer version of the Roman numeral is:", int_format)
}
