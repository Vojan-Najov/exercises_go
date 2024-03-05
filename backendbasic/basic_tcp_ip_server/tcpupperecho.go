// tcpupperecho serves tcp connections on port 8080, reading from each connection
// line-by-line and writting the upper-case version of each line back to the client.

package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "log"
  "net"
  "strings"
)

func main() {
  const name = "tcpupperecho"
  log.SetPrefix(name + "\t")

  // build the command-line interface
  port := flag.Int("p", 8080, "port to listen on")
  flag.Parse()

  // ListenTCP creates a TCP listener accepting connections on the  given address.
  // TCPAddr represent the address of a TCP end point; it has an IP, Port and Zone,
  // all of wich are optional
  // Zone only matters for IPv6, we'll ignore it for now.
  // If we omit the IP, it means we are listening on all available IP addresses;
  // if we omit the Port, it means we are listening on a random port.
  listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})
  if err != nil {
    panic(err)
  }
  defer listener.Close()
  log.Printf("listening at localhost: %s", listener.Addr())

  // loop forever, accepting connections one at a time
  for {
    // Accept() blocks until a connection is made, then returns a Conn representing
    // the connection
    conn, err := listener.Accept()
    if err != nil {
      panic(err)
    }
    // spawn a goroutine to handle the connection
    go echoUpper(conn, conn)
  }
}

// echoUpper reads lines from r, uppercases them, and writes them to w.
func echoUpper(w io.Writer, r io.Reader) {
  scanner := bufio.NewScanner(r)
  for scanner.Scan() {
    line := scanner.Text()
    log.Printf("received: %s", line)
    fmt.Fprintf(w, "%s\n", strings.ToUpper(line))
  }
  if err := scanner.Err(); err != nil {
    log.Printf("error: %s", err)
  }
}


