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
  timer := flag.Int("timer", 30, "Set your time")
  saveAnswer := flag.Bool("sA", false, "Save Your answer to CSV")

  flag.Parse()

  fmt.Printf("Timer set to %d \n", timer)
  go settimer(timer)



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


func settimer(timer *int){
  time.Sleep(time.Second)
  fmt.Println("You Have", *timer - 1, "left")
}