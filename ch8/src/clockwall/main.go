package main

import (
   "fmt"
   "bytes"
   "log"
   "net"
   "os"
   "strings"
)

func main() {
   for _, timeZone := range os.Args[1:] {
      request := strings.Split(timeZone, "=")
      timeZoneLocale := request[0]
      timeZoneServer := request[1]
      timeZoneHour := retrieveHour(timeZoneServer)
      fmt.Printf("%s %s", timeZoneLocale, timeZoneHour)
   }
}

func retrieveHour(timeZoneServer string) string {
   conn, err := net.Dial("tcp", timeZoneServer)
   if err != nil {
      log.Fatal(err)
      os.Exit(1)
   }
   defer conn.Close()
   var hour string
   buff := bytes.NewBufferString(hour)

   hourBytes := make([]byte, 9)
   _, err = conn.Read(hourBytes)
   if err != nil {
      log.Fatal(err)
      os.Exit(1)
   }

   _, err = buff.Write(hourBytes)
   if err != nil {
      log.Fatal(err)
      os.Exit(1)
   }

   return buff.String()
}

/* 
func mustCopy(dst io.Writer, src io.Reader) {
   if _, err := io.Copy(dst, src); err != nil {
      log.Fatal(err)
   }
}
*/
