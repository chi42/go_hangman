package main

import (
  "fmt"
  "os"
  "container/list"
//  "regexp"
//  "strings"
//  "rand"
//  "time"
)

const (
  BUF_SIZ = 100
  BLANK= '*'
)

var (
  word_list *list.List
)

// printers ...
func list_print(l *list.List) {

  for e := l.Front(); e != nil; e = e.Next() {
    fmt.Printf("%s\n", e)
  }
}


func main() {

  file_scan("dict")
//  rg,err := regexp.Compile("...")
}


// bring the whole dictionary into main memory
//    (based on the assumption that we will be guessing
//    at multiple words throughout the lifetime of the program)
func file_scan(dict_name string)  {
  word_list = new(list.List)
  f, err    := os.Open(dict_name)
  str       := ""
  num_bytes := 0
  var read_buf [BUF_SIZ]byte

  if f == nil {
    fmt.Printf("File error: %s\n", err.String())
    os.Exit(1)
  }

  num_bytes, err = f.Read(read_buf[:])

  // loop over file, until no more bytes to be read
  for ; num_bytes > 0; {
    // loop over buffer, newline indicates new word 
    for i := 0; i < num_bytes; i++ {
      if read_buf[i] == '\n' {
        b := []byte(str)
        word_list.PushBack(b)
        str = ""
      } else {
        str += string(read_buf[i])
      }
    }

    num_bytes, err = f.Read(read_buf[:])
  }
}


