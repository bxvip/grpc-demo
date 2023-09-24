package opc 

import (
	"fmt"
	"log"
	"net"
	"time"
)

type OpcSocket struct {
	conn      net.Conn
	sendChan  chan string
	recvChan  chan string
	closeChan chan struct{}
}

func NewOpcSocket(address string) (*OpcSocket, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to OPC server: %v", err)
	}

	socket := &OpcSocket{
		conn:      conn,
		sendChan:  make(chan string),
		recvChan:  make(chan string),
		closeChan: make(chan struct{}),
	}

	go socket.readLoop()
	go socket.writeLoop()

	return socket, nil
}

func (socket *OpcSocket) readLoop() {
	for {
		select {
		case <-socket.closeChan:
			return
		default:
			buf := make([]byte, 1024)
			n, err := socket.conn.Read(buf)
			if err != nil {
				log.Printf("Error reading from OPC server: %v", err)
				socket.Close()
				return
			}
			msg := string(buf[:n])
			socket.recvChan <- msg
		}
	}
}

func (socket *OpcSocket) writeLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-socket.closeChan:
			return
		case msg := <-socket.sendChan:
			_, err := socket.conn.Write([]byte(msg))
			if err != nil {
				log.Printf("Error writing to OPC server: %v", err)
				socket.Close()
				return
			}
		case <-ticker.C:
			socket.sendChan <- "Sample message" // 定时发送任务
		}
	}
}

func (socket *OpcSocket) Send(msg string) {
	socket.sendChan <- msg
}

func (socket *OpcSocket) Receive() string {
	return <-socket.recvChan
}

func (socket *OpcSocket) Close() {
	close(socket.closeChan)
	socket.conn.Close()
}
