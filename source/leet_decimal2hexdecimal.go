package main
import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
  "log"
  "math"
)
func toHex(num int) string {

  hex2dec_map := map[int]string{
    0:"0",
    1:"1",
    2:"2",
    3:"3",
    4:"4",
    5:"5",
    6:"6",
    7:"7",
    8:"8",
    9:"9",
    10:"a",
    11:"b",
    12:"c",
    13:"d",
    14:"e",
    15:"f",
  }
if num == 0{
 return "0"
}else if num > 0{
      remainders_slice:= make([]string,0,15)
  for{
    if num == 0{
      break
    }
    quotient := num/16
    fmt.Println("quotient is",quotient)
    remainder := num%16
    fmt.Println("remainder is",remainder)
    remainders_slice = append(remainders_slice,hex2dec_map[remainder])
    num = quotient
  }

  //fmt.Println("the unsorted integer slice is",remainders_slice)
  for i, j := 0, len(remainders_slice)-1; i < j; i, j = i+1, j-1 {
		remainders_slice[i], remainders_slice[j] = remainders_slice[j], remainders_slice[i]   // reverse the slice
	}
  //fmt.Println("the sorted slice is ",remainders_slice)

  return strings.Join(remainders_slice,"")
}else {
  num = num * (-1)
  remainders_slice_int:= make([]int,0,15)
  flipped_binary_slice := make([]int,0,15)
  hex_slice := make([]string,0,15)
  flipped_value := -1
  for{
    if num == 0{
      break
    }
    quotient := num/2
    //fmt.Println("quotient is",quotient)
    remainder := num%2
    //fmt.Println("remainder is",remainder)
    remainders_slice_int = append(remainders_slice_int,remainder)
    num = quotient
  }
  //fmt.Println("the unreversed remainder slice is",remainders_slice_int)
  if len(remainders_slice_int) < 32{
   leadiingzeros_count :=  32 - len(remainders_slice_int)
   for n := 0; n< leadiingzeros_count ;n++{
     remainders_slice_int = append(remainders_slice_int,0)
   }
  }
  //fmt.Println("the unreversed remainder slice with leading 0's is ",remainders_slice_int)
  for i, j := 0, len(remainders_slice_int)-1; i < j; i, j = i+1, j-1 {
    remainders_slice_int[i], remainders_slice_int[j] = remainders_slice_int[j], remainders_slice_int[i]   // reverse the slice
  }
  //fmt.Println("the binary equivalent of dec is   ",remainders_slice_int)

  for _, value := range remainders_slice_int{
    if value == 0{
      flipped_value = 1
    }else{
      flipped_value = 0
    }
    flipped_binary_slice = append(flipped_binary_slice,flipped_value)
}
//fmt.Println("the flipped binary is  ",flipped_binary_slice)
length :=len(flipped_binary_slice)
carry_over := 1
for x := (length -1); x > -1; x--{
  if (flipped_binary_slice[x] == 0 && carry_over ==1) {
      flipped_binary_slice[x] = 1
      carry_over = 0
      break
    }else if flipped_binary_slice[x] == 1 && carry_over == 1 {
      flipped_binary_slice[x] = 0
      carry_over = 1

      }
  }

//fmt.Println("the incremented binary is  ",flipped_binary_slice)
dec4hex := 0
a := 0
b := 0
c := 0
d := 0
m := length + 3
for {
  m = m - 4
  if m < 3{
    break
  }
  //fmt.Println("m is ***************** " ,m)
  a = flipped_binary_slice[m]* int(math.Pow(2, 0))
      //fmt.Println("m3 value is  : " ,flipped_binary_slice[m])
      //fmt.Println("a is : " ,a)
  b = flipped_binary_slice[m-1]* int(math.Pow(2, 1))
    //fmt.Println("m2 value is  : " ,flipped_binary_slice[m-1])
    //fmt.Println("b is : " ,b)
  c = flipped_binary_slice[m - 2]* int(math.Pow(2,2))
    //fmt.Println("m1 value is  : " ,flipped_binary_slice[m-2])
    //fmt.Println("c is : " ,c)
  d = flipped_binary_slice[m-3]* int(math.Pow(2,3))
    //fmt.Println("m0 value is  : " ,flipped_binary_slice[m-3])
    //fmt.Println("d is : " ,d)
  dec4hex = a + b + c + d
  //fmt.Println("first decimal is : " ,dec4hex)
  hex_slice = append(hex_slice,hex2dec_map[dec4hex])
}
for d, p := 0, len(hex_slice)-1; d < p; d, p = d+1, p-1 {
  hex_slice[d], hex_slice[p] = hex_slice[p], hex_slice[d]   // reverse the slice
}
hex_value := strings.Join(hex_slice,"")

return hex_value
}
}


func main(){

  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome ! This program converts integers to their hexadecimal equivalents.")
  fmt.Print("Please provide an Integer you would like to convert to a hexadecimal number : ")
  fmt.Print("\n")
  input, _ := consoleReader.ReadString('\n')

  upperString := strings.ToUpper(input)
  trimmedString := strings.TrimSpace(upperString)
  fmt.Println(trimmedString)

  int_value,e := strconv.Atoi(trimmedString)
  fmt.Println(int_value)
  if e == nil{
    str_hexadecimal := toHex(int_value)
    fmt.Println("The hexadcimal version of the integer is:", str_hexadecimal)
  } else{
    fmt.Println("This value",input," you entered is not an integer. Please try again and provide an integer.")
    log.Fatal(e)
  }

}
