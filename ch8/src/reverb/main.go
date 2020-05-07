package main

import (
   "net"
   "log"
   "time"
   "bufio"
   "fmt"
   "strings"
   "sync"
)

func main() {
   listener, err := net.Listen("tcp", "0.0.0.0:8000")
   if err != nil {
      log.Fatal(err)
   }
   for {
      conn, err := listener.Accept()
      if err != nil {
         log.Print(err)
         continue
      }
      handleCon(conn)
   }
}

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
   defer wg.Done()
   fmt.Fprintln(c, "\t", strings.ToUpper(shout))
   time.Sleep(delay)
   fmt.Fprintln(c, "\t", shout)
   time.Sleep(delay)
   fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func readInput(c net.Conn, input *bufio.Scanner, keepon chan<-struct{}) {
   var wg sync.WaitGroup
   for input.Scan() {
      keepon <- struct{}{}
      wg.Add(1)
      go echo(c, &wg, input.Text(), 1 * time.Second)
      wg.Wait()
   }
   log.Print("finish reading.")
}

func handleCon(c net.Conn) {
   defer c.Close()
   log.Print("starting handleCon...")
   input := bufio.NewScanner(c)

   keepon := make (chan struct{})
   go readInput(c, input, keepon)
   timeout := false
   for timeout == false {
      select {
      case <- time.After(10 * time.Second):
         log.Print("timeout reached.")
         timeout = true
      case <- keepon:
         log.Print("input read.")
         continue
      }
   }

   log.Print("closing connection")
}
