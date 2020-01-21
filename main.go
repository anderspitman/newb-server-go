package main

import (
        "log"
        "net"
        "os"
        "io"
        "strings"
)

func main() {
        log.Println("Starting up")

        ln, err := net.Listen("tcp", ":3838")
        if err != nil {
                // handle error
        }
        for {
                conn, err := ln.Accept()
                if err != nil {
                        log.Fatal(err)
                }
                go handleConnection(conn)
        }
}

func handleConnection(conn net.Conn) {

        defer conn.Close()

        var buf = make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
                log.Println(err)
        }

        if n == 0 {
                conn.Write([]byte("Must provide path\n"))
                return
        }

        path := strings.TrimSpace(string(buf[0:n]))

        if strings.Contains(path, "..") {
                conn.Write([]byte("Path cannot contain '..' \n"))
                return
        }

        relPath := strings.TrimLeft(path, "/")

        f, err := os.Open(relPath)
        if err != nil {
                log.Println(err)
                conn.Write([]byte("Not found \n"))
                return
        }
        defer f.Close()

        io.Copy(conn, f)

        //conn.Write([]byte(relPath))
}
