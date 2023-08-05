package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Item struct{
  Index int;
  Title string;
}

func greeting(){
  fmt.Println("Welcome to To Go")
  fmt.Println("a simple todo list tui application written in GoLang")
}

func printMenu() {
    fmt.Println("========= TODO LIST APP =========")
    fmt.Println("1. a to add new task")
//    fmt.Println("2. Toggle Task Done")
    fmt.Println("3. v to view all tasks")
    fmt.Println("4. ex to exit")
    fmt.Println("===============================")
}


func readItem() {
  fmt.Println("New Item: ")
  title, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Something when wrong", err)
  }

  item := Item{
    Index: itemIndex,
    Title: title,
  } 

  itemList = append(itemList, item)
  printListItem()
  itemIndex++
}

func printListItem(){
  for index, item := range itemList {
    if index + 1 == itemIndex{
      fmt.Print("x   ")
    }
    fmt.Print(item.Index)
    fmt.Print("   ")
    fmt.Print(item.Title)
  }
}

var reader = bufio.NewReader(os.Stdin)
var itemIndex = 1
var itemList []Item
func main() {
  greeting()
  printMenu()
  for {
    nav, _ := reader.ReadString('\n')
    nav = strings.TrimSpace(nav)
    switch nav {
      case "a":
        readItem()
      case "v":
        printListItem()
      case "ex":
        os.Exit(0)
      default:
        fmt.Println("we does not support that option")
    }
  
  }
  //var item string 
  //var itemlist []string
  //itemlist = append(itemlist, item)

  //fmt.Println(itemlist)
  // readItem(keyreader)
  
}



