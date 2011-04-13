package main

import (
  "fmt"
  "os"
  "container/list"
  "strings"
  "rand"
  "time"
)

const (
  BUF_SIZ = 100
  BLANK= '*'
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

// given a new word "word" we will now attempt to guess the word
// program will try until completion
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

  fmt.Printf("%2d %2d guess:  \tleft: %6d\t", total_tries, bad_tries, l_list_size)
  fmt.Printf("so far: %s\n", l_so_far)

  for {
    l, w := pick()

    lg, wg := try_guess(l, w)
    if !lg {
      bad_tries++
    }
    total_tries += 1
    if !updates (l, w) {
      fmt.Printf("No such word in dictionary!\n")
      return
    }

    fmt.Printf("%2d %2d guess: %c\tleft: %6d\t", total_tries, bad_tries, l, l_list_size)
    fmt.Printf("so far: %s\n", l_so_far)

    if wg {
      //fmt.Printf("<%d %s\n>%d %s\n", total_tries, l_word, bad_tries, l_word)
      break
    }
  }
  fmt.Printf("\n")

}

// based on counters and current progress, attempt to guess a new
// letter or word
func pick () (byte, string) {

  // a pseudo randomish seed, not perfect but better then
  // an obviously deterministic random number generater
  rand.Seed(time.Nanoseconds())

  max_val   := uint(0)
  max_pos   := uint(0)

  // pick the "obvious" matches, i.e., certain letters
  // are the only possible letters that can fit in a spot
  // so we pick those first
  for i, _ := range l_counts.pos {
    for j, _ := range l_counts.pos[i] {

      if l_counts.pos[i][j] == uint(l_list_size) {
        max_pos = uint(j + 65)
        //fmt.Printf("OBVIOUS!!!!\n")

        break
      } else {
        if l_counts.pos[i][j] != 0 {
          break
        }
      }
    }

    if max_pos > 0 {
      break
    }
  }

  // weren't able to pick an obvious match, so we pick the
  // letter that occurs in the most words
  if max_pos == 0 {
    max_total   := uint(0)
    mod         := 10.0

    for i, v := range l_counts.uniq {
      if v > max_val {
        max_val   = v
        max_pos   = uint(i)
        max_total = l_counts.total[i]
      }

      // the tie breakers
      if v == max_val {
         // pick the letter randomly
        if rand.Int() % int(mod) == 0 {
          //fmt.Printf("***RAND SELECTED\n")
          max_val   = v
          max_pos   = uint(i)
          max_total = l_counts.total[i]

          // this number is choosen somewhat arbitrarily
          // the intention is to randomly select the new letter
          // with decreasing probability 
          mod += (float64(i) * 2.0) 

        // pick the letter that occurs overall the most
        } else if l_counts.total[i] > max_total {
          //fmt.Printf("***SWITCH MADE\n")
          max_val = v
          max_pos = uint(i)
          max_total = l_counts.total[i]
        }
      }
    }
    max_pos += 65
  }

  return byte(max_pos), ""
}


// update counts and various other variables after a letter
// or word is guessed
func updates (l byte, w string) bool {

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

          if l_list_size == 0 {
            return false
          }

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

  return true
}

// update the counters when a single word 'w' is stripped 
func word_removal_count(w string) {
  var temp [26]uint
  var index int

  // iterate across all letters in w
  // for each letter update the counters
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

// evaluate the status of the guess that uses
// either the character 'l' or the word 'w'
// return:
//    success of letter guess, overall success (i.e if word is completed)
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


// generate counters initially
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

// initial build, scan the dictionary and bring into 
// main memory all the words of matching length
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

