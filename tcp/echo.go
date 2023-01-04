package tcp

import (
	"bufio"
	"context"
	"fmt"
	"godis/lib/sync/wait"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Bool
}

func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	err := c.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	client := &EchoClient{
		Conn: conn,
	}
	h.activeConn.Store(client, struct{}{})

	reader := bufio.NewReader(conn)
	for {
		if h.closing.Load() {
			fmt.Println("close...")
			_ = conn.Close()
			return
		}
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
				h.activeConn.Delete(client)
			} else {
				log.Println(err)
			}
		}

		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

func (h *EchoHandler) Close() error {
	log.Println("handler shutting down...")
	h.closing.Store(true)
	h.activeConn.Range(func(key, value any) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true
	})
	return nil
}

// MakeEchoHandler creates EchoHandler
func MakeEchoHandler() *EchoHandler {
	return &EchoHandler{}
}
