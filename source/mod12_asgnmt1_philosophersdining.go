package main
import (
  "fmt"
 "sync"
)

type Chopstick struct {
  sync.Mutex
}

type Philosopher struct{
  id int
  leftCS, rightCS *Chopstick
}

var currently_eating chan int

func (p Philosopher) eat(wg1 *sync.WaitGroup){
  var wg2 sync.WaitGroup

  for count := 1; count < 4; count++{
    wg2.Add(1)
    go host(p,&wg2)
    wg2.Wait()
    p.leftCS.Lock()
    p.rightCS.Lock()
    fmt.Println("starting to eat <",p.id," > \n")

    p.leftCS.Unlock()
    p.rightCS.Unlock()
    fmt.Println("finishing eating <",p.id,"> \n")
    //fmt.Println("\n")
     <- currently_eating
    //fmt.Print("diner who just got out ", a)
  //  fmt.Println("\n")
  }
  wg1.Done()
}

func host(requesting_philosopher Philosopher,wg3 *sync.WaitGroup){
  currently_eating <- requesting_philosopher.id
  wg3.Done()
  //fmt.Println("current diners are")
}

func main(){
  Chopsticks := make([]*Chopstick,5)
  for i := 0;i < 5; i ++{
    //new returns pointert to/address of  the Chopstick
    Chopsticks[i] = new(Chopstick)
  }

  philosophers := make([]*Philosopher,5)
  for n := 0; n < 5;n ++{
    philosophers[n] = &Philosopher{n+1,Chopsticks[n],Chopsticks[(n + 1)%5]}
    //fmt.Println(philosophers[n].id,n,(n + 1)%5)

  }

  var wg sync.WaitGroup
  wg.Add(5)
  currently_eating = make(chan int,2)

  for x := 0; x < 5; x++{
    go philosophers[x].eat(&wg)
  }

  wg.Wait()
}
