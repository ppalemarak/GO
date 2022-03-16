package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Animal interface {
Eat()
Move()
Speak()
}

type Cow struct{
food,locomotion,noise string
}
func (cow_a Cow) Eat(){
  fmt.Println(cow_a.food)
}
func (cow_b Cow) Move(){
  fmt.Println(cow_b.locomotion)
}

func (cow_c Cow) Speak(){
  fmt.Println(cow_c.noise)
}

type Bird struct{
food,locomotion,noise string
}

func (bird_a Bird) Eat(){
  fmt.Println(bird_a.food)
}
func (bird_b Bird) Move(){
  fmt.Println(bird_b.locomotion)
}

func (bird_c Bird) Speak(){
  fmt.Println(bird_c.noise)
}


type Snake struct{
food,locomotion,noise string
}
//snake := Animal{"mice","slither","hsss"}
func (snake_a Snake) Eat(){
  fmt.Println(snake_a.food)
}
func (snake_b Snake) Move(){
  fmt.Println(snake_b.locomotion)
}

func (snake_c Snake) Speak(){
  fmt.Println(snake_c.noise)
}


func main(){

fmt.Print("Welcome ! This program will help you create  3 types of animals - cow, bird, snake!\n")
fmt.Print("Each of these animals will be created with these 3 properties -  what they eat, how they move and what they speak.\n")
fmt.Print("You can  pick a name for your animal and indicate which animal you want to create by entering 3 keywords in this format 'newanimal <name you woukd like to use for the animal> <animal type>'.e.g. Enter 'newanimal Jack cow' to create a cow named Jack.\n")
fmt.Print("You can indicate what you want to learn about the animals you created by entering 3  keywords in this format 'query <name of animal> <animal action>.e.g. Enter 'query Jack eat' to learn what the cow named Jack eats.\n")
fmt.Print("Allowed keywords for  <animal type> are 'cow,bird,snake'.\n")
fmt.Print("Allowed keywords for  <animal action> are 'eat,speak,move'.\n")
fmt.Print("Animal name is case insenstive'.\n")


created_animals := make(map[string]Animal)

 for{
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("\n")
  fmt.Print("Please indicate if you would like to create or learn about an animal by entering in the format described above . Or, enter X to exit .")
  fmt.Print("\n")
  fmt.Print(">")
  input, _ := consoleReader.ReadString('\n')
  lowerString :=  strings.ToLower(input)
  trimmedString := strings.TrimSpace(lowerString)

  if trimmedString == "x"{
    break
  }

  strings_slice := strings.Split(trimmedString," ")

  command_word := strings_slice[0]
  animal_name := strings_slice[1]
  command_parameter := strings_slice[2]

  animal_details,animal_created := created_animals[animal_name]

  if command_word == "newanimal"{

   if !animal_created{
      switch command_parameter{
      case "cow":
        created_animals[animal_name] = Cow{"grass", "walk", "moo"}
        fmt.Print("Created it!\n")
      case "bird":
        created_animals[animal_name] = Bird{"worms", "fly", "peep"}
        fmt.Print("Created it!\n")
      case "snake":
        created_animals[animal_name] =  Snake{"mice","slither","hsss"}
        fmt.Print("Created it!\n")
      default:
        fmt.Print(command_parameter, " is not an animal that can be created by this program.Please try again and specify if you would like to create a cow,bird or snake.\n")
      }
      } else{
        fmt.Print(animal_name," is an animal name that has already been used.Please use a new name.\n")
      }

  }else if command_word == "query"{

      if animal_created {
      switch command_parameter{
      case "eat":
        animal_details.Eat()
      case "move":
        animal_details.Move()
      case "speak":
        animal_details.Speak()
      default:
        fmt.Print(command_parameter, " is not an animal action that is known for this animal.Please try again and specify a valid animal action.\n")
      }
    }else{
      fmt.Print(animal_name, " is an animal name that has not been created yet .Please try again and use  an animal name you previously created to query.\n")
    }


  }else{
    fmt.Print(command_word, " is not a valid keyword.Please try again and enter a valid keyword.\n")
    continue
  }


}

}
