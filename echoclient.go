package main

import (
        "log"
        "net"
)

func main() {
        var conns []net.Conn
        for i := 0; i < 20000; i++ {
                c, err := net.Dial("tcp", "", "10.0.0.15:2020")
                if err != nil {
                        log.Fatalf("dial error: %v", err)
                }
                conns = append(conns, c)
        }
        select {}
}
