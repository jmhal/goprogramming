package main

import (
   "fmt"
   "bufio"
   "os"
   "net"
   "log"
   "strings"
   "io"
   "strconv"
)

const BUFFERSIZE = 1024

func main() {
   server := os.Args[1]
   port := os.Args[2]
   conn, err := net.Dial("tcp", server + ":" + port)
   if err != nil {
      log.Fatal(err)
   }
   defer conn.Close()

   consoleReader := bufio.NewReader(os.Stdin)
   connReader := bufio.NewReader(conn)
   connWriter := bufio.NewWriter(conn)

   fmt.Printf("Connecting %s:%s\n", server, port)
   for {
      fmt.Print("> ")
      cmd, _ := consoleReader.ReadString('\n')
      cmd = strings.TrimSpace(cmd)
      cmdSlice := strings.Split(cmd, " ")

      switch cmdSlice[0] {
         case "ls":
            if _, err := connWriter.WriteString("ls\n"); err == nil {
               err = connWriter.Flush()
               if err != nil {
                  log.Fatal(err)
                  os.Exit(1)
               }
               files, nerr := connReader.ReadString('\n')
               if nerr != nil {
                  log.Fatal(nerr)
               }
               fmt.Println(files)
               continue
            }
            os.Exit(1)
         case "cd":
            if _, err := connWriter.WriteString("cd\n"); err == nil {
               err = connWriter.Flush()
               if err != nil {
                  log.Fatal(err)
                  os.Exit(1)
               }
               if _, err := connWriter.WriteString(cmdSlice[1] + "\n"); err == nil {
                  err = connWriter.Flush()
                  if err != nil {
                     log.Fatal(err)
                     os.Exit(1)
                  }
               }
               continue
            }
            os.Exit(1)
         case "get":
            if _, err := connWriter.WriteString("get\n"); err != nil {
               log.Printf("Problem sending the command get %s:", err)
               os.Exit(1)
            }
            if _, err := connWriter.WriteString(cmdSlice[1] + "\n"); err != nil {
               log.Printf("Problem sending the get argument: %s", err)
               os.Exit(1)
            }

            if err := connWriter.Flush(); err != nil {
               log.Printf("Problem flush: %s", err)
               os.Exit(1)
            }


            fileName := "./" + cmdSlice[1]
            if file, err := os.Create(fileName); err == nil {
               defer file.Close()

               bufferFileSize := make([]byte, 10)
               conn.Read(bufferFileSize)
               fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

               var receivedBytes int64
               for {
                  if (fileSize - receivedBytes) < BUFFERSIZE {
                     io.CopyN(file, conn, (fileSize - receivedBytes))
			            conn.Read(make([]byte, (receivedBytes + BUFFERSIZE) - fileSize))
			            break
		            }
		            io.CopyN(file, conn, BUFFERSIZE)
                  receivedBytes += BUFFERSIZE
               }
            } else {
               log.Println(err)
               os.Exit(1)
            }
            continue
         case "put":
            conn.Write([]byte("put\n"))

            fileName := cmdSlice[2]
            stat, err := os.Stat(cmdSlice[1])

            if err != nil || stat.IsDir() {
               log.Println(err)
               os.Exit(1)
            }

            fileName = fillString(fileName, 64)
            conn.Write([]byte(fileName))

            fileSize := fillString(strconv.FormatInt(stat.Size(), 10), 10)
            conn.Write([]byte(fileSize))

            if file, err := os.Open(cmdSlice[1]); err == nil {
               defer file.Close()
               sendBuffer := make([]byte, BUFFERSIZE)
               for {
                  _, err := file.Read(sendBuffer)
                  if err == io.EOF {
                     break
                  }
                  conn.Write(sendBuffer)
               }
            } else {
               log.Println(err)
               os.Exit(1)
            }

            continue
         case "close":
            if _, err := connWriter.WriteString("close\n"); err == nil {
               err = connWriter.Flush()
               if err != nil {
                  log.Fatal(err)
                  return 
               }
            }
            return
      }
   }
}

func mustCopy(dst io.Writer, src io.Reader) {
   if _, err := io.Copy(dst, src); err != nil {
      log.Fatal(err)
   }
}

func fillString(returnString string, toLength int) string {
	for {
		lengtString := len(returnString)
		if lengtString < toLength {
			returnString = returnString + ":"
			continue
		}
		break
	}
	return returnString
}
