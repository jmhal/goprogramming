package main

import (
   "net"
   "log"
   "io"
   "io/ioutil"
   "os"
   "strings"
   "bufio"
   "strconv"
)

const BUFFERSIZE = 1024

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

func handleCon(c net.Conn) {
   defer c.Close()
   pwd := "/"
   cb := bufio.NewReader(c)
   for {
      // Extract Request
      operation, err := cb.ReadString('\n')
      if err != nil {
         log.Printf("Invalid Request: %s", err)
         break
      }
      // Dispatcher
      switch strings.TrimSpace(operation) {
         case "ls" :
            var files []string
            fileInfo, err := ioutil.ReadDir(pwd)
            if err != nil {
               log.Print(err)
               c.Write([]byte("Invalid directory.\n"))
               continue
            }
            for _, file := range fileInfo {
               if file.IsDir() {
                  files = append(files, file.Name() + "/")
               } else {
                  files = append(files, file.Name())
               }
            }
            c.Write([]byte(strings.Join(files[:]," ")))
            c.Write([]byte("\n"))
         case "cd" :
            parameter := extractParameter(cb)
            log.Printf("cd parameter: %s", parameter)

            pwd = setupPath(pwd, parameter)
            log.Println(pwd)
            continue
         case "get" :
            parameter := extractParameter(cb)
            log.Printf("get parameter: %s", parameter)

            filePath := setupPath(pwd, parameter)
            stat, err := os.Stat(filePath)
            if err == nil && !stat.IsDir() {
               log.Printf("file to send: %s", filePath)
            } else {
               log.Println(err)
            }

            fileSize := fillString(strconv.FormatInt(stat.Size(), 10), 10)
            log.Printf("file size: %s", fileSize)
            c.Write([]byte(fileSize))

            if file, err := os.Open(filePath); err == nil {
               defer file.Close()
               sendBuffer := make([]byte, BUFFERSIZE)
               for {
                  _, err := file.Read(sendBuffer)
                  if err == io.EOF {
                     break
                  }
                  c.Write(sendBuffer)
               }
            } else {
               log.Println(err)
            }

            continue
         case "put" :
            bufferFileName := make([]byte, 64)
            c.Read(bufferFileName)
            fileName := strings.Trim(string(bufferFileName), ":")
            fileName = setupPath(pwd, fileName)

            bufferFileSize := make([]byte, 10)
            c.Read(bufferFileSize)
            fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

            if file, err := os.Create(fileName); err == nil {
               defer file.Close()

               var receivedBytes int64
               for {
                  if (fileSize - receivedBytes) < BUFFERSIZE {
                     io.CopyN(file, c, (fileSize - receivedBytes))
			            c.Read(make([]byte, (receivedBytes + BUFFERSIZE) - fileSize))
			            break
		            }
		            io.CopyN(file, c, BUFFERSIZE)
                  receivedBytes += BUFFERSIZE
               }
            } else {
               log.Println(err)
               continue 
            }

            continue
         case "close":
            return
      }
   }
}

func setupPath(pwd string, parameter string) string {
   if strings.HasPrefix(parameter, "/") {
      return parameter
   } else {
      if pwd == "/" {
         return pwd + parameter
      } else {
         return pwd + "/" + parameter
      }
   }
}

func extractParameter(cb *bufio.Reader) string {
   parameter, err := cb.ReadString('\n')
   if err != nil {
      log.Print(err)
   }
   return strings.TrimSpace(parameter)
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
