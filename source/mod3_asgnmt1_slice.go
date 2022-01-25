package main
import (
  "fmt"
  "sort"
  "strconv"
  //"bufio"
  //"os"
)


func main(){
  var x string
  int_slice := make([]int,0,3)
  var userInput bool = true

  //input := bufio.NewScanner(os.Stdin)
	//for input.Scan()
  for userInput{
      fmt.Printf("Please enter the integer you would like to add to the list . Enter X if you are done : ")
      fmt.Scan(&x)
      if x == "X"{
        //userInput = false
        break
      }
      y,e := strconv.Atoi(x)
      if e == nil{
        int_slice  =  append(int_slice,y)
        sort.Ints(int_slice)
        fmt.Println("The list of integers you entered in ascending order are :", int_slice)
      }else{
        fmt.Println("It looks like you didn't enter an integer. Please check and enter an integer.")
      }

  }


  fmt.Println("Thank you!")

}
