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

  lis := file_scan("/home/chi/code/go_hangman/dict", 18)

  for e := lis.Front(); e != nil; e = e.Next() {
    fmt.Printf("%s\n", e.Value)
  }
}

func char_count(lis *list.List) []uint {
  var counter [26]uint

  for i := 0; i < 26; i++ {
    counter[i] = 0
  }

  for e := lis.Front(); e != nil; e.Next() {
    for _, val :=  range e.Value {
      switch val {
        case 'A' :
          counter[A] += 1
        case 'B' :
          counter[B] += 1
        case 'C' :
          counter[C] += 1
        case 'D' :
          counter[D] += 1
        case 'E' :
          counter[E] += 1
        case 'F' :
          counter[F] += 1
        case 'G' :
          counter[G] += 1
        case 'H' :
          counter[H] += 1
        case 'I' :
          counter[I] += 1
        case 'J' :
          counter[J] += 1
        case 'K' :
          counter[K] += 1
        case 'L' :
          counter[L] += 1
        case 'M' :
          counter[M] += 1
        case 'N' :
          counter[N] += 1
        case 'O' :
          counter[O] += 1
        case 'P' :
          counter[P] += 1
        case 'Q' :
          counter[Q] += 1
        case 'R' :
          counter[R] += 1
        case 'S' :
          counter[S] += 1
        case 'T' :
          counter[T] += 1
        case 'U' :
          counter[U] += 1
        case 'V' :
          counter[V] += 1
        case 'W' :
          counter[W] += 1
        case 'X' :
          counter[X] += 1
        case 'Y' :
          counter[Y] += 1
        case 'Z' :
          counter[Z] += 1
      }
    }
  }

  return counter[:]
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

