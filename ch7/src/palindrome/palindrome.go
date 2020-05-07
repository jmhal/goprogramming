package main

import (
   "fmt"
   "sort"
)

type Palindrome []rune

func (p Palindrome) Len() int { return len(p)}
func (p Palindrome) Less(i, j int) bool { return p[i] < p[j] }
func (p Palindrome) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func isPalindrome(s sort.Interface) bool {
   tam := s.Len()
   for i, j := 0,  tam - 1 ; i < tam / 2; i, j = i + 1, j - 1 {
     if !(!s.Less(i,j) && !s.Less(j,i)) {
        return false
     }
   }
   return true
}

func main() {
   var a Palindrome
   a = []rune("João Marcelo")
   fmt.Println(isPalindrome(a))
   a = []rune("123321")
   fmt.Println(isPalindrome(a))
   a = []rune("1235321")
   fmt.Println(isPalindrome(a))
   a = []rune("arara")
   fmt.Println(isPalindrome(a))
   a = []rune("banana")
   fmt.Println(isPalindrome(a))
}
