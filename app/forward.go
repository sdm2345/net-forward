package app

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func StartForward(listens, rules []string) {

	ctx, cancel := context.WithCancel(context.Background())

	for i, rule := range rules {
		go ListenTcp(ctx, listens[i], rule)
	}
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	cancel()
	os.Exit(0)

}

func ListenTcp(ctx context.Context, fromAddr, toAddr string) {

	log.Println("start from", fromAddr, "to", toAddr)
	conn, err := net.Listen("tcp", fromAddr)
	if err != nil {
		log.Println("error", err)
		return
	}
	log.Println("start from", fromAddr, "to", toAddr, "ok")
	go func() {
		select {
		case <-ctx.Done():
			log.Println("done")
			_ = conn.Close()
		}
	}()
	for {
		conn, err := conn.Accept()
		if err != nil {
			continue
		}
		//建立 一个远程连接
		client, err := net.Dial("tcp", toAddr)
		if err != nil {
			log.Println("err", err)
			_ = conn.Close()
			continue
		}
		log.Println("connect from", conn.RemoteAddr())
		closeConn := func() {
			log.Println("close", conn.RemoteAddr())
			_ = conn.Close()
			_ = client.Close()
		}
		go func() {
			defer closeConn()
			_, _ = io.Copy(client, conn)
		}()
		go func() {
			defer closeConn()
			_, _ = io.Copy(conn, client)
		}()

	}
}
