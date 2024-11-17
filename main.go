package main

import (
    "flag"
    "fmt"
)

func main() {
    mode := flag.String("mode", "client", "Mode client/rendezvous")
    ip := flag.String("ip", "localhost", "IP address of the rendezvous server")

    flag.Parse()

    switch *mode {
    case "client":
        StartClient(*ip)
    case "rendezvous":
        StartRendezvousServer()
    default:
        fmt.Println("Illegal mode")
    }
}
