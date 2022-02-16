package main
import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
  "log"
)

func Swap(swap_slice []int, position int){
  x := swap_slice[position]
  swap_slice[position] = swap_slice[position + 1]
  swap_slice[position + 1 ] = x
}

func Bubblesort(sort_slice []int){
  for{
      var swapped bool = false
      for index,_ := range sort_slice{
        if index < (len(sort_slice) -1 ){
          //fmt.Println("in bubblesote")
          if sort_slice[index] > sort_slice[index + 1]{
            Swap(sort_slice, index )
            swapped = true
            //fmt.Println("still iterating")
          }
        }

      }
      if swapped == false{
        break
      }

  }
}

func main(){

  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome ! This program sorts integers in ascending order.")
  fmt.Print("Please provide a list of integers you would like to sort separated by comma. You can provide up to 10 integers. e.g. 40,-100,1,-90 :  ")
  fmt.Print("\n")
  input, _ := consoleReader.ReadString('\n')
  //fmt.Println(input)
  trimmedString := strings.TrimSpace(input)
	string_slice := strings.Split(trimmedString, ",")
  size_list := len (string_slice)
  //fmt.Println("in string format ", string_slice)
  //fmt.Println("size i ", size_list)
  integer_slice := make([]int,0,size_list)

  for _, value := range string_slice{

    //fmt.Println("current value is ",value)

    /*if i == (size_list - 1){
      fmt.Println(" breaking out of for ")
      break
    }*/
    int_value,e := strconv.Atoi(value)

    //fmt.Println(" where is it breaking")
    if e == nil{
      integer_slice = append(integer_slice,int_value)
      //fmt.Println(input_slice)
    } else{
      fmt.Println("This element ",value," is not an integer. Please try again and ensure your list contains only integers.")
      log.Fatal(e)
    }

  }
//fmt.Println("The unsorted list is ", integer_slice)
Bubblesort(integer_slice)
fmt.Println("The sorted order is :", integer_slice)
}
