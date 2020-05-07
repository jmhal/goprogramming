package main

import (
  "fmt"
  "bufio"
  "os"
  "io"
  "bytes"
)

type ByteCounter int
type WordCounter int
type LineCounter int

type Counter struct {
   internalWriter io.Writer
   nbytes *int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
   *c += ByteCounter(len(p))
   return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
   scanner := bufio.NewScanner(bytes.NewReader(p))
   scanner.Split(bufio.ScanWords)
   count := 0
   for scanner.Scan() {
      count += 1
   }
   *c += WordCounter(count)
   if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "lendo entrada: ", err)
   }
   return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
   scanner := bufio.NewScanner(bytes.NewReader(p))
   scanner.Split(bufio.ScanLines)
   count := 0
   for scanner.Scan() {
      count += 1
   }
   *c += LineCounter(count)
   if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "lendo entrada: ", err)
   }
   return count, nil
}

func (c Counter) Write(p []byte) (int, error) {
   *(c.nbytes) += int64(len(p))
   return c.internalWriter.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
   var c Counter
   var i int64
   c.nbytes = &i
   c.internalWriter = w
   return c, c.nbytes
}

func main() {
   var c ByteCounter
   c.Write([]byte("hello"))
   fmt.Println(c)

   var a WordCounter
   a.Write([]byte("casa comida roupa lavada amor"))
   fmt.Println(a)

   var b LineCounter
   b.Write([]byte("casa comida\n roupa\n lavada\n"))
   fmt.Println(b)

   d, n := CountingWriter(os.Stdout)
   d.Write([]byte("João Marcelo Uchôa de Alencar\n"))
   fmt.Println(*n)
   d.Write([]byte("abc\n"))
   fmt.Println(*n)
}
