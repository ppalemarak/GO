package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "log"
)

func main(){
    type Person struct{
      fname string
      lname string
    }
    var persons_slice = make([]Person,0,3)
    var filename string
    consoleReader := bufio.NewReader(os.Stdin)

    fmt.Print("what is the name of the text file that you want to read from? ")
    input, _ := consoleReader.ReadString('\n')
    filename = strings.TrimSpace(input)

    file_handle, err := os.Open(filename)
    if err != nil {
    fmt.Print("Please make sure your text file is present in the same directory as the source GO file ")
    fmt.Print("\n")
		log.Fatal(err)
    //fmt.Print("Please make sure your text file is present in the same directory as the source GO file")
  }

    scanner := bufio.NewScanner(file_handle)
  	for scanner.Scan() {
  		s := strings.Split(scanner.Text(), " ")
  		var person_x Person
      if len(s[0]) > 20 {
      s[0] = string(s[0][:20])
      }
      if len(s[1]) > 20 {
      s[1] = string(s[1][:20])
      }
  		person_x.fname, person_x.lname = s[0], s[1]
  		persons_slice = append(persons_slice, person_x)

  	}
    file_handle.Close()

    fmt.Print("The first and last names  listed in your file '",filename,"' are :")
    fmt.Print("\n")
    for _, v := range persons_slice {
    		fmt.Println(v.fname,v.lname)
    	}





}
