package main

import (
   "fmt"
   "net"
   "log"
   "bufio"
   "time"
)

type client struct {
   channel chan<- string
   name string
}

var (
   entering = make(chan client)
   leaving = make (chan client)
   messages = make(chan string)
)

func main() {
   listener, err := net.Listen("tcp", "localhost:8000")
   if err != nil {
      log.Fatal(err)
   }

   go broadcaster()

   for {
      conn, err := listener.Accept()
      if err != nil {
         log.Print(err)
         continue
      }
      go handleConn(conn)
   }
}

func handleConn(conn net.Conn) {
   input := bufio.NewScanner(conn)

   login := make(chan string)
   go func(login chan<- string, input *bufio.Scanner) {
      for input.Scan() {
         login <- input.Text()
         return
      }
   }(login, input)

   who := <-login
   log.Println(who)

   ch := make(chan string)
   go clientWriter(conn, ch)

   //who := conn.RemoteAddr().String()

   ch <- "You are " + who
   messages <- who + " has arrived"
   cl := client {
      channel: ch,
      name: who,
   }
   entering <- cl

   action := make(chan bool)
   go func(action <-chan bool, conn net.Conn, who string) {
      for {
         timeout := time.Tick(300 * time.Second)
         select {
         case a := <-action:
               if a == false {
                  return;
               }
         case <-timeout:
               log.Println("Removing idle user: " + who)
               conn.Close()
               return;
         }
      }
   } (action, conn, who)

   for input.Scan() {
      messages <- who + ": " + input.Text()
      action <- true
   }

   leaving <- cl
   action <- false
   messages <- who + " has left"
   conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
   for msg := range ch {
      fmt.Fprintln(conn, msg)
   }
}

func broadcaster() {
   clients := make(map[client]bool)
   for {
      select {
      case msg := <-messages:
         for cli := range clients {
            timeout := time.Tick(5 * time.Second)
            select {
            case cli.channel <- msg:
               continue;
            case <- timeout:
               continue;
            }
         }
      case cli := <-entering:
         clients[cli] = true
         onlineusers := fmt.Sprintf("Online users: \n")
         for cl := range clients {
            onlineusers = fmt.Sprintf("%s %s\n", onlineusers, cl.name)
         }
         fmt.Println(onlineusers)
         for cl := range clients {
            cl.channel <- onlineusers
         }
      case cli := <-leaving:
         delete(clients, cli)
         close(cli.channel)
      }
   }
}
