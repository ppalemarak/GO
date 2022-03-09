package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Animal struct {
food,locomotion,noise string

}

func (animal_a Animal) Eat(){
  fmt.Println(animal_a.food)
}

func (animal_b Animal) Move(){
  fmt.Println(animal_b.locomotion)
}

func (animal_c Animal) Speak(){
  fmt.Println(animal_c.noise)
}


func main(){
  cow := Animal{"grass", "walk", "moo"}
  bird := Animal{"worms", "fly", "peep"}
  snake := Animal{"mice","slither","hsss"}

fmt.Print("Welcome ! This program will help you learn about 3 types of animals - cow, bird, snake! \n")
fmt.Print("Things you can learn about these animals - what they eat, how they move and what they speak.\n")
fmt.Print("You can indicate what you want to learn by entering two keywords in this format  <animal type> <animal action>.e.g. Enter 'cow speak' to learn what the cow speaks. \n ")
fmt.Print("Allowed keywords for  <animal type> are 'cow,bird,snake'.\n")
fmt.Print("Allowed keywords for  <animal action> are 'eat,speak,move'. \n")

 for{
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("\n")
  fmt.Print("Please let us know what would you like to learn in the format <animal type> <animal action> . Or, enter X to exit .")
  fmt.Print("\n")
  fmt.Print(">")
  input, _ := consoleReader.ReadString('\n')
  lowerString := strings.ToLower(input)
  trimmedString := strings.TrimSpace(lowerString)
  if trimmedString == "x"{
    break
  }

  strings_slice := strings.Split(trimmedString," ")

/*  if strings_slice[1] == "eat"{
    strings_slice[0].Eat()
  }else if strings_slice[1] == "speak"{

  }else if strings_slice[1] == "move"{

  }else{
    fmt.Print(strings_slice[1], " is not a valid keyword.Please try again and enter a valid keyword")
    continue
  }*/

if strings_slice[0] == "cow"{
  if strings_slice[1] == "eat"{
    cow.Eat()
  }else if strings_slice[1] == "speak"{
    cow.Speak()
  }else if strings_slice[1] == "move"{
    cow.Move()
  }else{
    fmt.Print(strings_slice[1], " is not a valid keyword.Please try again and enter a valid keyword.\n")
    continue
  }
}else if strings_slice[0] == "bird"{
  if strings_slice[1] == "eat"{
    bird.Eat()
  }else if strings_slice[1] == "speak"{
    bird.Speak()
  }else if strings_slice[1] == "move"{
    bird.Move()
  }else{
    fmt.Print(strings_slice[1], " is not a valid keyword.Please try again and enter a valid keyword.\n")
    continue
  }
}else if strings_slice[0] == "snake"{
  if strings_slice[1] == "eat"{
    snake.Eat()
  }else if strings_slice[1] == "speak"{
    snake.Speak()
  }else if strings_slice[1] == "move"{
    snake.Move()
  }else{
    fmt.Print(strings_slice[1], " is not a valid keyword.Please try again and enter a valid keyword.\n")
    continue
  }
}else{
  fmt.Print(strings_slice[0], " is not a valid keyword.Please try again and enter a valid keyword.\n")
  continue
}


}

}
