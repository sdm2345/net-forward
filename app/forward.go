package app

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func StartForward(rules []string) {

	ctx, cancel := context.WithCancel(context.Background())

	for _, rule := range rules {

		go Listen(ctx, rule)
	}
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("\n- Ctrl+C pressed in Terminal")
	cancel()
	os.Exit(0)

}

func Listen(ctx context.Context, rule string) {

	rule = strings.Trim(rule, "/")
	arr := strings.Split(rule, "/")
	if len(arr) == 0 {
		panic(errors.New("error params:" + rule))
	}

	//todo 先支持 tcp
	if arr[0] == "tcp" {
		_, _, err := net.SplitHostPort(strings.Join(arr[1:3], ":"))
		if err != nil {
			panic(err)
		}
		_, _, err = net.SplitHostPort(strings.Join(arr[4:], ":"))
		if err != nil {
			panic(err)
		}

		ListenTcp(ctx, strings.Join(arr[1:3], ":"), strings.Join(arr[4:], ":"))
	} else {
		panic(errors.New("error params:" + rule))
	}
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
