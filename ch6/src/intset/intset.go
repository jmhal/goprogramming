package intset

import (
   "fmt"
   "bytes"
)

type IntSet struct {
   words []uint64
}

func (s *IntSet) Has(x int) bool {
   word, bit := x/64, uint(x % 64)
   return word < len(s.words) && s.words[word] & (1<<bit) != 0
}

func (s *IntSet) Add(x int) {
   word, bit := x/64, uint(x % 64)
   for word >= len(s.words) {
      s.words = append(s.words, 0)
   }
   s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(vals...int) {
   for _, v := range vals {
      s.Add(v)
   }
}

func (s *IntSet) UnionWith(t *IntSet) {
   for i, tword := range t.words {
      if i < len(s.words) {
         s.words[i] |= tword
      } else {
         s.words = append(s.words, tword)
      }
   }
}

func (s *IntSet) IntersectWith(t *IntSet) {
   for i, tword := range t.words {
      if i < len(s.words) {
         s.words[i] &= tword
      }
   }
}

func (s *IntSet) DifferenceWith(t *IntSet) {
   u := s.Copy()
   u.IntersectWith(t)
   for i, w := range u.words {
      s.words[i] = s.words[i] ^ w
   }
}

func (s *IntSet) SymmetricDifference(t *IntSet) (*IntSet) {
   v := s.Copy()
   v.DifferenceWith(t)
   x := t.Copy()
   x.DifferenceWith(s)
   v.UnionWith(x)
   return v
}

func (s *IntSet) Len() int {
   var length uint64
   for _, w := range s.words {
      for w > 0 {
         length += w & 1
         w >>= 1
      }
   }
   return int(length)
}

func (s *IntSet) Remove(x int) {
   word, bit := x/64, uint(x % 64)
   if (word < len(s.words) && s.words[word] & (1<<bit) != 0) {
      s.words[word] = s.words[word] ^ 1 << bit
   }
}

func (s *IntSet) Clear() {
   for i := 0; i < len(s.words); i++ {
      s.words[i] = 0
   }
}

func (s *IntSet) Copy() *IntSet {
   var newset IntSet
   for i := 0; i < len(s.words); i++ {
      newset.words = append(newset.words, s.words[i])
   }
   return &newset
}

func (s *IntSet) Elems() []int {
   var elems []int
   for i, word := range s.words {
      if word == 0 {
      }
      for j := 0; j < 64; j++ {
         if word & (1<<uint(j)) != 0 {
	    elems = append(elems, 64*i+j)
	 }
      }
   }
   return elems
}

func (s *IntSet) String() string {
   var buf bytes.Buffer
   buf.WriteByte('{')
   for i, word := range s.words {
      if word == 0 {
         continue
      }
      for j := 0; j < 64; j++ {
         if word & (1<<uint(j)) != 0 {
	    if buf.Len() > len("{") {
	       buf.WriteByte(' ')
	    }
	    fmt.Fprintf(&buf, "%d", 64*i+j)
	 }
      }
   }
   buf.WriteByte('}')
   return buf.String()
}
