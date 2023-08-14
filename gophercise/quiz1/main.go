package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main()  {

  var newfilename string
  flag.StringVar(&newfilename, "rn", "", "Provide a new name")
  saveAnswer := flag.Bool("sA", false, "Save Your answer to CSV")  
  flag.Parse()

    bufio.NewReader(os.Stdin)
  file,err := os.Open("problem.csv")
  if err != nil {
    fmt.Println(err)
  }

  filereader := csv.NewReader(file)
  records,_ := filereader.ReadAll()

  if *&newfilename != ""{
    fmt.Println("New File Name: ",*&newfilename)
  }

  if *saveAnswer{
    fmt.Println("Save Answer to new file")
  }
  csreader := bufio.NewReader(os.Stdin)
  for _,record := range records{
    fmt.Println(record[0])
  
    value,err := csreader.ReadString('\n')
    if err != nil {
      fmt.Println(err)
    } 

    if *(&value) == record[1]{
      fmt.Println("correct")
    }else {
      fmt.Println("false")
    }
  }
}
