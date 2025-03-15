package main
import (
  "net"
  "fmt"
  "flag"
  "log"
)

func main (){
   server := flag.String("s","","server")
   porta := flag.String("p","","porta")

   flag.Parse()

   strServer := *server
   strPorta := *porta

   if strServer == "" && strPorta == "" {
        fmt.Println("Potevi scrivere -s server -p porta , comunque... ")
        fmt.Println(" ")
        fmt.Print("Server: ")
        fmt.Scanln(&strServer)
        fmt.Print("Porta: ")
        fmt.Scanln(&strPorta)
   }

   _, err := net.Dial("tcp", strServer + ":" + strPorta)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println("Risultato: OK")
}
