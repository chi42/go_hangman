package main

import (
  "fmt"
  "os"
  "container/list"
  "strings"
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

var (
  l_lis *list.List
  l_counts counters
  l_so_far []byte
  l_word string

  l_list_size int
)



type counters struct {
  uniq []uint
  total []uint
  pos [][26]uint

}


func main() {

  if len(os.Args) > 1 {
    try_word(strings.ToUpper(os.Args[1]))
  }

}

func try_word (word string) {
  total_tries := 0
  bad_tries := 0
  l_word        = word
  word_len      := len(l_word)
  l_so_far      = make([]byte, word_len)

  for i,_ := range l_so_far {
    l_so_far[i] = BLANK
  }

  file_scan("dict", word_len)
  char_count(word_len)

  for {
    l, w := pick()
    lg, wg := try_guess(l, w)
    if !lg {
      bad_tries++
    }
    total_tries += 1
    fmt.Printf("%d %d guess: %c\tleft: %d\t\t", total_tries, bad_tries, l, l_list_size)

    // if word all filled in, then quit

    updates (l, w)
    fmt.Printf("so far: %s\n", l_so_far)

    if wg {
      break
    }
  }
  fmt.Printf("\n")

}



func pick () (byte, string) {

  max_val   := uint(0)
  max_pos   := uint(0)

  for i, v := range l_counts.uniq {
    if v > max_val {
      max_val   = v
      max_pos   = uint(i)
    }
  }

  return byte(max_pos + 65), ""
}



func updates (l byte, w string) {

  if l > 0 {
    index := l - 65
    l_counts.total[index] = 0
    l_counts.uniq[index]  = 0
    for i, _ := range l_counts.pos {
      l_counts.pos[i][index] = 0
    }

    var e_prev *list.Element

    e_prev = l_lis.Front().Next()
    for e := l_lis.Front(); e != nil; e = e.Next() {
      for i, val := range e.Value.(string) {
        // two kinds of words to remove:
        //    word does not contain the guessed letter at the same spot(s)
        //    or word contains the guessed letter, but not in the same spot(s)
        if (l_so_far[i] != BLANK && l_so_far[i] != uint8(val)) ||
              (l_so_far[i] != l && uint8(val) == l) {
          word_removal_count(e.Value.(string))
          l_list_size--
          l_lis.Remove(e)
          e = e_prev
          break
        }
      }
      e_prev = e
    }
  // consider the instance of where we guessed a wrong, and guessed wrong
  } else {

  }
}

func word_removal_count(w string) {
  var temp [26]uint
  var index int

  for pos, v := range w {
    index = v - 65
    temp[index] = 1
    if l_counts.total[index] > 0 {
      l_counts.total[index] -= 1
    }
    if l_counts.pos[pos][index] > 0 {
      l_counts.pos[pos][index] -=1
    }
  }

  for i, v := range temp {
    if l_counts.uniq[i] > 0 {
     l_counts.uniq[i] -= v
    }
  }

}

func try_guess (l byte, w string) (bool, bool) {
  l_match := false

  if l > 0 {
    for i,v := range l_word {
      if byte(v) == l {
        l_match = true
        l_so_far[i] = l
      }
    }
  }

 for _, v := range l_so_far {
   if v == BLANK {
     return l_match, false
   }
 }

 return true, true
}

func char_count(word_len int) {
  var temp [26]uint

  l_counts.uniq     = make([]uint, 26)
  l_counts.total    = make([]uint, 26)
  l_counts.pos      = make([][26]uint, word_len)

  // for each word
  for e := l_lis.Front(); e != nil; e = e.Next() {

    // (string) needed, as type assertion for interface type
    // for each letter in word
    for i, val := range e.Value.(string) {
      l_counts.total[val - 65] += 1
      temp[val - 65] = 1
      l_counts.pos[i][val-65] += 1
    }

    // set and init
    for i, v := range temp {
      l_counts.uniq[i] += v
      temp[i] = 0
    }
  }

}


func file_scan(name string, word_len int)  {
  l_list_size = 0
  l_lis     = new(list.List)
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
          l_list_size++
          l_lis.PushBack(str)
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

}

