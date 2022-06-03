package main
import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
  "log"
  "sort"
 "sync"
)

var string_slice []string
var size_list int

func sort_slice_1(slice_1 []int, wg *sync.WaitGroup){
  fmt.Print("1st goroutine is sorting this subarray : ", slice_1,"\n")
  sort.Ints(slice_1)
  //fmt.Print("1st partition sorted: ", slice_1,"\n")
  wg.Done()
}

func sort_slice_2(slice_2 []int, wg *sync.WaitGroup){
  fmt.Print("2nd goroutine is sorting this subarray : ", slice_2,"\n")
  sort.Ints(slice_2)
  //fmt.Print("2nd partition sorted: ", slice_2,"\n")
  wg.Done()
}

func sort_slice_3(slice_3 []int, wg *sync.WaitGroup){
  fmt.Print("3rd goroutine is sorting this subarray : ", slice_3,"\n")
  sort.Ints(slice_3)
  //fmt.Print("3rd partition sorted: ", slice_3,"\n")
  wg.Done()
}

func sort_slice_4(slice_4 []int, wg *sync.WaitGroup){
  fmt.Print("4th goroutine is sorting this subarray : ", slice_4,"\n")
  sort.Ints(slice_4)
  //fmt.Print("4th partition sorted: ", slice_4,"\n")
  wg.Done()
}

func main(){
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome ! This program sorts integers in ascending order.\n")
  for{
    fmt.Print("Please provide a list of integers you would like to sort separated by comma. Please  provide at least 8 integers. e.g. 40,-100,1,-90 :  ")
    fmt.Print("\n")
    input, _ := consoleReader.ReadString('\n')

    trimmedString := strings.TrimSpace(input)
  	string_slice = strings.Split(trimmedString, ",")
    size_list = len (string_slice)
    if size_list < 8{
      fmt.Println("The number of integers you have provided is less than 8. Please try again and provide at least 8 integers!")
      continue
    }else{
      break
    }
  }

  /*fmt.Println("in string format ", string_slice)*/
  //fmt.Println("size is ", size_list)

  integer_slice := make([]int,0,size_list)

  for _, value := range string_slice{
    int_value,e := strconv.Atoi(value)
    if e == nil{
      integer_slice = append(integer_slice,int_value)
    } else{
      fmt.Println("This element ",value," is not an integer.Please try again and ensure your list contains only integers.")
      log.Fatal(e)
    }

  }
  //fmt.Println("The unsorted order is :", integer_slice)

  sub_size,extra_size := size_list/4,size_list%4

  var integer_slice_1,integer_slice_2,integer_slice_3,integer_slice_4 []int
  integer_slice_1 = integer_slice[0:sub_size]
  //fmt.Println("integer_slice_1 is:", integer_slice_1,len(integer_slice_1),cap(integer_slice_1))

  switch extra_size{
    case 0:
      integer_slice_2 = integer_slice[sub_size:2*sub_size]
      //fmt.Println("integer_slice_2 is:", integer_slice_2,len(integer_slice_2),cap(integer_slice_2))

      integer_slice_3 = integer_slice[2*sub_size:3*sub_size]
      //fmt.Println("integer_slice_3 is:", integer_slice_3,len(integer_slice_3),cap(integer_slice_3))

      integer_slice_4 = integer_slice[3*sub_size:]
      //fmt.Println("integer_slice_4 is:", integer_slice_4,len(integer_slice_4),cap(integer_slice_4))

    case 1:
      integer_slice_2 = integer_slice[sub_size:2*sub_size]
      //fmt.Println("integer_slice_2 is:", integer_slice_2,len(integer_slice_2),cap(integer_slice_2))

      integer_slice_3 = integer_slice[2*sub_size:3*sub_size]
      //fmt.Println("integer_slice_3 is:", integer_slice_3,len(integer_slice_3),cap(integer_slice_3))

      integer_slice_4 = integer_slice[3*sub_size:]
      //fmt.Println("integer_slice_4 is:", integer_slice_4,len(integer_slice_4),cap(integer_slice_4))

    case 2:
      integer_slice_2 = integer_slice[sub_size:2*sub_size]
      //fmt.Println("integer_slice_2 is:", integer_slice_2,len(integer_slice_2),cap(integer_slice_2))

      integer_slice_3 = integer_slice[2*sub_size:(3*sub_size)+ 1]
      //fmt.Println("integer_slice_3 is:", integer_slice_3,len(integer_slice_3),cap(integer_slice_3))

      integer_slice_4 = integer_slice[((3*sub_size)+ 1):]
      //fmt.Println("integer_slice_4 is:", integer_slice_4,len(integer_slice_4),cap(integer_slice_4))

    case 3:
      integer_slice_2 = integer_slice[sub_size:(2*sub_size) + 1]
      //fmt.Println("integer_slice_2 is:", integer_slice_2,len(integer_slice_2),cap(integer_slice_2))

      integer_slice_3 = integer_slice[((2*sub_size) + 1):(3*sub_size)+ 2]
      //fmt.Println("integer_slice_3 is:", integer_slice_3,len(integer_slice_3),cap(integer_slice_3))

      integer_slice_4 = integer_slice[(3*sub_size+ 2):]
      //fmt.Println("integer_slice_4 is:", integer_slice_4,len(integer_slice_4),cap(integer_slice_4))
  }

 //fmt.Println("integer_slice_main is:", integer_slice)

  var wg sync.WaitGroup
  wg.Add(4)
  go sort_slice_1(integer_slice_1, &wg)
  go sort_slice_2(integer_slice_2, &wg)
  go sort_slice_3(integer_slice_3, &wg)
  go sort_slice_4(integer_slice_4, &wg)
  wg.Wait()

  sorted_integer_slice := []int{}
  sorted_integer_slice = append(integer_slice_1,integer_slice_2...)
  sorted_integer_slice = append(sorted_integer_slice,integer_slice_3...)
  sorted_integer_slice = append(sorted_integer_slice,integer_slice_4...)

  sort.Ints(sorted_integer_slice)
  fmt.Println("Sorted list:", sorted_integer_slice)
}
