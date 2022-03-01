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

func GenDisplaceFn(acceleration,initial_velocity,initial_displacement float64) func (float64) float64{
  fn := func(time float64) float64{
    return (0.5*acceleration*(math.Pow(time, 2)) + initial_velocity*time + initial_displacement)

  }
  return fn
}


func main(){

  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome ! This program  calculates displacement as a function of time, acceleration , initial velocity and initial displacement.")
  fmt.Print("Please provide the values(float) for acceleration,initial velocity and initial displacement separated by comma.Items after the 3rd item will be ignored :  ")
  fmt.Print("\n")
  input, _ := consoleReader.ReadString('\n')
  //fmt.Println(input)
  trimmedString := strings.TrimSpace(input)
  strings_slice := strings.Split(trimmedString, ",")
  size_list := len (strings_slice)

  float_slice := make([]float64,0,size_list)

  for i := 0; i < 3; i++{
    float_value,e := strconv.ParseFloat(strings_slice[i],64)

    if e == nil{
      float_slice = append(float_slice,float_value)
    } else{
      fmt.Println("This item ",strings_slice[i]," is not a float. Please try again and ensure your list contains only floats.")
      log.Fatal(e)
    }

  }

  displacementFn := GenDisplaceFn(float_slice[0],float_slice[1],float_slice[2])

  fmt.Print("Now please provide the value(float) for time : ")
  fmt.Print("\n")
  input2, _ := consoleReader.ReadString('\n')

  trimmedString2 := strings.TrimSpace(input2)
  time_float,e2 := strconv.ParseFloat(trimmedString2,64)

  if e2 == nil{
    displacement := displacementFn(time_float)
    fmt.Println("The calculated displacement is ",displacement," for these inputs :")
    fmt.Println( "Acceleration = ",float_slice[0],",initial velocity = ",float_slice[1],",initial displacement = ", float_slice[2], ",time =",time_float)
  }else{
   fmt.Println("The time entered ",trimmedString2," is not a float. Please try again and ensure you provide a float value for time.")
   log.Fatal(e2)
 }



}
