package popcount

// pc [i] é a população de i
var pc [256]byte

func init() {
   for i := range pc {
      pc[i] = pc[i/2] + byte(i&1)
   }
}

// PopCount devolve a população (número d ebits definidos) de x
func PopCount(x uint64) int {
   return int(pc[byte(x>>(0*8))] +
      pc[byte(x>>(1*8))] +
      pc[byte(x>>(2*8))] +
      pc[byte(x>>(3*8))] +
      pc[byte(x>>(4*8))] +
      pc[byte(x>>(5*8))] +
      pc[byte(x>>(6*8))] +
      pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
   var value byte
   for i := 0; i < 8; i++ {
      value += pc[byte(x>>(uint(i)*8))]
   }
   return int(value)
}

func PopCount64(x uint64) int {
  numbersOne := 0
  for i := 0; i < 64; i++ {
     if (x & 1) == 1 {
        numbersOne++
     }
     x = x >> 1
  }
  return numbersOne
}

func PopCountClean(x uint64) int {
  numbersOne := 0
  for x != 0 {
     x = x & (x - 1)
     numbersOne++
  }
  return numbersOne
}
