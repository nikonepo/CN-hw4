package main

import (
    "fmt"
    "net"
    "time"
)

func StartClient(ip string) {
    serverAddr, err := net.ResolveUDPAddr("udp", ip+":12345")
    if err != nil {
        panic(err)
    }

    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        panic(err)
    }

    defer conn.Close()

    _, err = conn.Write([]byte("Hello from client"))
    if err != nil {
        fmt.Printf("Couldn't send the message: %v\n", err)
        return
    }

    buf := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buf)
    if err != nil {
        fmt.Printf("Couldn't read the message: %v\n", err)
        return
    }

    peerAddr, err := net.ResolveUDPAddr("udp", string(buf[:n]))
    if err != nil {
        fmt.Printf("Couldn't resolve peer address: %v\n", err)
        return
    }

    peerConn, err := net.ListenUDP("udp", nil)
    if err != nil {
        fmt.Printf("Couldn't create UDP connection: %v\n", err)
        return
    }

    defer peerConn.Close()

    fmt.Printf("Received: %s\n", string(buf[:n]))

    go listen(conn)

    for {
        _, err := peerConn.WriteToUDP([]byte("Message from client"), peerAddr)
        if err != nil {
            fmt.Printf("Couldn't send the message: %v\n", err)
            return
        }

        fmt.Printf("Sent message to peer\n")
        time.Sleep(time.Second * 5)
    }
}

func listen(conn *net.UDPConn) {
    buf := make([]byte, 1024)
    for {
        n, addr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Printf("Couldn't read the message: %v\n", err)
            return
        }

        fmt.Printf("Received from %s: %s\n", addr.String(), string(buf[:n]))
    }
}
