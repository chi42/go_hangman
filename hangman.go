package main

import (
  "fmt"
  "os"
  "container/list"
)

const (
  BUF_SIZ = 100
  BLANK= '*'

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

type counters struct {
  uniq []uint
  total []uint
  pos [][26]uint

}


func main() {

  //lis := file_scan("dict", 6)

  //for e := lis.Front(); e != nil; e = e.Next() {
  //  fmt.Printf("%s\n", e.Value)
  //}
  //counts := char_count(6, lis)

  ////try_word("hello")

  //for i, _ := range counts.total {
  //  //fmt.Printf("** %c %d\n", i + 65, val)
  //  fmt.Printf("*  %c %d\n", i + 65, counts.uniq[i])
  //}

  //fmt.Printf("\n\n")
  //for i, v1 := range counts.pos {
  //  for j := 0; j < 26; j++ {
  //    fmt.Printf("%d ", v1[j])
  //  }
  // fmt.Printf(" %d\n", i)
  //}

  //l, _ := pick(lis, counts, 10)
  //fmt.Printf("PICKED: %d, %c\n\n\n", l, l + 65)


  ////update_count(lis, counts, nil, 'A', "")
  //for i, _ := range counts.total {
  //  //fmt.Printf("** %c %d\n", i + 65, val)
  //  fmt.Printf("*  %c %d\n", i + 65, counts.total[i])
  //}

  //for i, _ := range counts.pos {
  //  for j := 0; j < 26; j++ {
  //    fmt.Printf("%d ", counts.pos[i][j])
  //  }
  // fmt.Printf(" %d\n", i)
  //}

  try_word("ADIDAS")

}

func try_word (word string) {
  tries_left    := 6
  word_len      := len(word)
  so_far        := make([]byte, word_len)

  for i,_ := range so_far {
    so_far[i] = BLANK
  }

  lis := file_scan("dict", word_len)
  counts := char_count(word_len, lis)

  for {
    l, w := pick(lis, counts, tries_left)
    fmt.Printf("guess: %c\n", l)

    // if word all filled in, then quit
    if try_guess(word, so_far, l, w) {
      break
    }

    tries_left--

    fmt.Printf("so far: %s\n", so_far)
  }

}



func pick (lis *list.List, counts *counters, tries_left int) (byte, string) {

  max_val   := uint(0)
  max_pos   := uint(0)

  for i, v := range counts.total {
    if v > max_val {
      max_val   = v
      max_pos   = uint(i)
    }
  }

  counts.total[max_pos] = 0
  counts.uniq[max_pos]  = 0
  for i, _ := range counts.pos {
    counts.pos[i][max_pos] = 0
  }

  return byte(max_pos + 65), ""
}



func update_count (lis *list.List, counts *counters, so_far []byte,
  l byte, w string) {

  var temp [26]uint

  if l > 0 {
    for e := lis.Front(); e != nil; e = e.Next() {
      for i, val := range e.Value.(string) {
        if uint8(val) == l {
          if counts.total[val - 65] > 0 {
            counts.total[val - 65] -= 1
          }
          if counts.pos[i][val - 65] > 0 {
            counts.pos[i][val - 65] -= 1
          }
          temp[val - 65] = 1
        }
      }

      for i, v := range temp {
        if counts.uniq[i] > 0 {
          counts.uniq[i] -= v
        }
        temp[i] = 0
      }
    }
  } else {

  }

}


func try_guess (word string, so_far []byte, l byte, w string) bool {

  if l > 0 {
    for i,v := range word {
      if byte(v) == l {
        so_far[i] = l
      }
    }
  }

 for _, v := range so_far {
   if v == BLANK {
     return false
   }
 }

 return true
}

func char_count(word_len int, lis *list.List) (*counters) {
  var temp [26]uint

  counts         := new(counters)
  counts.uniq     = make([]uint, 26)
  counts.total    = make([]uint, 26)
  counts.pos      = make([][26]uint, word_len)

  // for each word
  for e := lis.Front(); e != nil; e = e.Next() {

    // (string) needed, as type assertion for interface type
    // for each letter in word
    for i, val := range e.Value.(string) {
      counts.total[val - 65] += 1
      temp[val - 65] = 1
      counts.pos[i][val-65] += 1
    }

    // set and init
    for i, v := range temp {
      counts.uniq[i] += v
      temp[i] = 0
    }
  }

  return counts
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

