// writetcp connect to a TCP server at localhost with specified port (8080 by default)
// and forwards stdin to the server, line-by-line, until EOF is reached.
// Received lines from the server are printed to stdout.

package main

import (
  "bufio"
  "flag"
  "fmt"
  "log"
  "net"
  "os"
)

func main() {
  const name = "writetcp"
  log.SetPrefix(name + "\t")

  // register the command-line flags: -p specified the port to connect to
  port := flag.Int("p", 8080, "port to connect to")
  flag.Parse()

  conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})
  if err != nil {
    log.Fatalf("error connecting to localhost:%d: %v", *port, err)
  }
  log.Printf("connect to %s: will forward stdin", conn.RemoteAddr())
  defer conn.Close()

  // spawn a goroutine to read incoming lines from the server and print them to stdout
  // TCP is full-duplex, so we can read and write at the same time;
  // we just need to spawn a goroutine to do the reading
  go func() {
    for connScanner := bufio.NewScanner(conn); connScanner.Scan(); {
      fmt.Printf("%s\n", connScanner.Text())
      if err := connScanner.Err(); err != nil {
        log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
      }
    }
  }()

  // read incoming lines from stdin and forward them to the server.
  for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); {
    log.Printf("sent: %s\n", stdinScanner.Text())
    if _, err := conn.Write(stdinScanner.Bytes()); err != nil {
      log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
    }
    if _, err := conn.Write([]byte("\n")); err != nil {
      log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
    }
    if stdinScanner.Err() != nil {
      log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
    }
  }

}
