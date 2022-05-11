package main
import (
  "fmt"
  "time"
)

func AssignAndIncrement(integer_address_a *int){
  //Step (a):initialze an integer to 1
  *integer_address_a = 1

  //Step (b):increment it by 1
  *integer_address_a = *integer_address_a + 1
}

func Display(integer_address_b *int){
  fmt.Println("*************************************************************************")
  //Step (c):print the incremented value
  fmt.Println("The incremented value of the integer is " ,*integer_address_b)
  fmt.Println("*************************************************************************")
  fmt.Println("\n")
}

func main(){
  fmt.Print("\n---------------------------------------------------------------WELCOME!-----------------------------------------------------------------------------------------------------------\n")
  fmt.Println("The goal of this program is to demonstrate the problem called Race Condition that can occur in GO when goroutines need to communicate with each other .")
  fmt.Println("\n")
  fmt.Print("Expected goal for this program is to :\n")
  fmt.Print("Step (a):initialze an integer to 1\n")
  fmt.Print("Step (b):increment it by 1  \n")
  fmt.Print("Step (c):print the incremented value.\n")
  fmt.Print("And the expected outome is 2.")
  fmt.Println("\n")
  fmt.Println("However if you run this program multiple times you will notice that the program is not behaving as expected.You may get different outputs when you run this program at different times - sometimes 2 and sometimes 0.")
  fmt.Println("What you are observing and experiencing is the Race Condition problem!")
  fmt.Print("This Race Condition is caused because the program implemented the steps (a),(b) and (c) in two separate goroutines that need to communicate with each other, are accessing a shared memory location (address of the integer) and are dependent on the non-deterministic ordering of interleaving .\n")
  fmt.Print("goroutine 'AssignAndIncrement': implements steps (a) and (b),\n")
  fmt.Print("groutine 'Display': implements step(c)\n")
  fmt.Print("The goroutine 'Display' is expected to print the incremented value(step c) and is  dependent on goroutine 'AssignAndIncrement' which initializes(step a) and increments the integer(step c). Hence in the program the goroutine 'Display' is scheduled to execute after the goroutine 'AssignAndIncrement' but since these are gorutines they are run concurrently and the individual instructions/steps within these goroutines are interleaved at the machine code level. The ordering of interleaving is NON-DETERMINISTIC and is handled at the machine code level and is completely out of control of the GO source code. This is why Steps b and c are  being interleaved in different sequences at the machine code level each time the prorgram is run. Ideally step (c) should be executed after step (b) but since these steps are part of separate goroutines which are running concurrently they are being sequenced in a random order at the machine code level \n")
  fmt.Println("\n")
  fmt.Println("Please run the program at least 5 or 6 times to observe the program provide different outputs -  either 2 or 0 ")
  integer_address := new(int)
  go AssignAndIncrement(integer_address)
  go Display(integer_address)
  time.Sleep(1 * time.Second)
  fmt.Print("---------------------------------------------------------------THANK YOU!---------------------------------------------------------------------------------------------------------------\n")
}
