package main

import(
  "fmt"
  server "api/server"
)
func main() {
  fmt.Println("In main..Going to call server")
server.ServerStart()
}
