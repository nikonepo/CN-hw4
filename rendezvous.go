package main

import (
    "fmt"
    "net"
)

type ClientInfo struct {
    Addr *net.UDPAddr
}

var clients []ClientInfo

func StartRendezvousServer() {

    addr, err := net.ResolveUDPAddr("udp", ":12345")
    if err != nil {
        panic(err)
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        panic(err)
    }

    defer conn.Close()

    fmt.Printf("Rendezvous server started on %s\n", conn.LocalAddr().String())

    for {
        buf := make([]byte, 1024)
        n, addr, err := conn.ReadFromUDP(buf)

        if err != nil {
            fmt.Println("Error: ", err)
            continue
        }

        if len(clients) < 2 {
            clients = append(clients, ClientInfo{addr})
            fmt.Printf("Client %d connected from %s\n", len(clients), addr.String())
        }

        fmt.Printf("Client %d connected from %s\n", len(clients), addr.String())

        if len(clients) == 2 {
            for _, client := range clients {
                otherClient := clients[1]
                if client.Addr.String() == addr.String() {
                    otherClient = clients[0]
                }

                message := fmt.Sprintf("%s", otherClient.Addr.String())
                conn.WriteToUDP([]byte(message), client.Addr)
            }

            clients = []ClientInfo{}
        }

        fmt.Printf("Received %d bytes: %s\n", n, string(buf))
    }
}
