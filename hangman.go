package main

import (
  "fmt"
  "os"
  "container/list"
)

const (
  BUF_SIZ = 100

  // leters
  A = 0
  B = 1
  C = 2
  D = 3
  E = 4
  F = 5
  G = 6
  H = 7
  I = 8
  J = 9
  K = 10
  L = 11
  M = 12
  N = 13
  O = 14
  P = 15
  Q = 16
  R = 17
  S = 18
  T = 19
  U = 20
  V = 21
  W = 22
  X = 23
  Y = 24
  Z = 25
)

func main() {

  lis := file_scan("/home/chi/code/go_hangman/dict", 6)

  //for e := lis.Front(); e != nil; e = e.Next() {
  //  fmt.Printf("%s\n", e.Value)
  //}
  //uniq_count, total_count, pos_count := char_count(18, lis)
  _, _, pos_count := char_count(6, lis)

  //for i, val := range total_count {
  //  fmt.Printf("** %c %d\n", i + 65, val)
  //  fmt.Printf("*  %c %d\n", i + 65, uniq_count[i])
  //}

  //fmt.Printf("\n\n")
  //for _, i := range pos_count {
  //  for j := 0; j < 26; j++ {
  //    fmt.Printf("%d ", i[j])
  //  }
  // fmt.Printf("\n")
  //}

}

func char_count(word_len int, lis *list.List) ([]uint, []uint, [][26]uint) {
  var total_count, uniq_count, temp [26]uint
  pos_count := make([][26]uint, word_len)

  // for each word
  for e := lis.Front(); e != nil; e = e.Next() {

    // (string) needed, as type assertion for interface type
    // for each letter in word
    for i, val :=  range e.Value.(string) {
      total_count[val - 65] += 1
      temp[val - 65] = 1
      pos_count[i][val-65] += 1
    }

    // set and init
    for i, v := range temp {
      uniq_count[i] += v
      temp[i] = 0
    }
  }

  return uniq_count[:], total_count[:], pos_count[:]
}

func file_scan(name string, word_len int)  *list.List {
  lis       := new(list.List)
  f, err    := os.Open(name, os.O_RDONLY, 0666)
  str       := ""
  num_b     := 0
  i         := 0
  var store_a [BUF_SIZ]byte

  if f == nil {
    fmt.Printf("File error: %s\n", err.String())
    os.Exit(1)
  }

  // loop over file, until no more bytes to be read
  num_b, err = f.Read(store_a[:])
  for ; num_b > 0; {
    // iterate through byte array, break if end of line 
    for ; i < num_b; i++ {
      if store_a[i] == '\n' {
        if len(str) == word_len {
          lis.PushBack(str)
        }
        str = ""
        i++
        break
      }
      str += string(store_a[i])
    }

    // read more bytes if no more bytes in buffer
    if i == num_b {
      num_b, err = f.Read(store_a[:])
      i = 0
    }
  }

  return lis
}

