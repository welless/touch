package main

import (
  "net"
  "fmt"
  "os"
)

func main() {
  ns, err := net.LookupHost("gz-cdb-nvmjxb1y.sql.tencentcdb.com")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
    return
  }

  for _, n := range ns {
    fmt.Fprintf(os.Stdout, "--%s\n", n)
  }
  
}

