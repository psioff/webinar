package main

import (
    "fmt"
    "net"
    "os"
)

const N = 0xffff - (20+8)

func main() {
    addr, err := net.ResolveUDPAddr("udp",":7")
    if err  != nil {
	fmt.Println("error: ",err)
	os.Exit(0)
    }
    c, err := net.ListenUDP("udp", addr)
    if err  != nil {
	fmt.Println("error: ",err)
	os.Exit(0)
    }
    defer c.Close()

    buf := make([]byte, N)
    for {
        n, ra, err := c.ReadFromUDP(buf)
	fmt.Println(n)
        c.WriteToUDP(buf[:n], ra)
        if err != nil {
            fmt.Println("error: ", err)
        }
    }
}
