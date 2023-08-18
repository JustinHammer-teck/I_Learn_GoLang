package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {

  var newfilename string
  flag.StringVar(&newfilename, "rn", "", "Provide a new name")
  timer := flag.Int("timer", 5, "Set your time")
  saveAnswer := flag.Bool("sA", false, "Save Your answer to CSV")

  flag.Parse()

  timerup := make(chan bool, 1)

  fmt.Printf("Timer set to %d \n", *timer)
  go settimer(timer, timerup)

  bufio.NewReader(os.Stdin)

  file,err := os.Open("problem.csv")
  if err != nil {
    fmt.Println(err)
  }

  if *&newfilename != ""{
    fmt.Println("New File Name: ",*&newfilename)
  }

  if *saveAnswer{
    fmt.Println("Save Answer to new file")
  }

  filereader := csv.NewReader(file)
  records,_ := filereader.ReadAll()
  csreader := bufio.NewReader(os.Stdin)

  for {

    select {
    case <- timerup:
      os.Exit(0)
    default:
      for _,row := range records{
        fmt.Println( "How much is", row[0])

        value, err := csreader.ReadString('\n')
        if err != nil{
          fmt.Println(err)
        }

        value = strings.TrimSpace(value)

        fmt.Println(value == row[1])
      }
    }

  }
}


func settimer(timer *int, timeup chan bool){
  ref := *timer

  for i := ref; i > 0; i-- {
    time.Sleep(time.Second)

    fmt.Println(i)
  }

  timeup <- true
  close(timeup)
}