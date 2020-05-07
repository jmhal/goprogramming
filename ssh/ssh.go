// Origem:
// http://blog.ralch.com/tutorial/golang-ssh-connection/
package main

import (
   "fmt"
   "os"
   "io"
   "io/ioutil"
   "golang.org/x/crypto/ssh"
)

func PublicKeyFile(file string) ssh.AuthMethod {
   buffer, err := ioutil.ReadFile(file)
   if err != nil {
      return nil
   }

   key, err := ssh.ParsePrivateKey(buffer)
   if err != nil {
      return nil
   }
   return ssh.PublicKeys(key)
}

func main() {
   username := os.Args[1]
   keyfile := os.Args[2]
   hostport := os.Args[3]

   sshConfig := &ssh.ClientConfig{
      User: username,
      HostKeyCallback: ssh.InsecureIgnoreHostKey(),
      Auth: []ssh.AuthMethod{
         PublicKeyFile(keyfile)},
   }

   connection, err := ssh.Dial("tcp", hostport, sshConfig)
   if err != nil {
      fmt.Printf("Failed to dial: %s\n", err)
      return
   }
   fmt.Println("Conexão feita.")

   session, err := connection.NewSession()
   if err != nil {
      fmt.Printf("Failed to create session: %s\n", err)
      return
   }
   fmt.Println("Sessão criada.")

   modes := ssh.TerminalModes {
      ssh.ECHO:          0,     // disable echoing
      ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
      ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
   }

   if err := session.RequestPty("xterm", 120, 80, modes); err != nil {
      session.Close()
      fmt.Errorf("request for pseudo terminal failed: %s", err)
      return
   }

   stdin, err := session.StdinPipe()
   if err != nil {
      fmt.Errorf("Unable to setup stdin for session: %v", err)
      return
   }
   go io.Copy(stdin, os.Stdin)

   stdout, err := session.StdoutPipe()
   if err != nil {
      fmt.Errorf("Unable to setup stdout for session: %v", err)
      return
   }
   go io.Copy(os.Stdout, stdout)

   stderr, err := session.StderrPipe()
   if err != nil {
      fmt.Errorf("Unable to setup stderr for session: %v", err)
      return
   }
   go io.Copy(os.Stderr, stderr)

   err = session.Run("w")

   return
}
