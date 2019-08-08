package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var filecontent = "hello the file streamer worked correctly"

func TCPFileReceiver(file string) (err error) {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return fmt.Errorf("can't listen on port : %v", err)
	}
	f, _ := os.Create(file)
	defer f.Close()
	fmt.Println("in file receiver, no actions yet")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		fmt.Println("got to right before filestream")
		go fileStream(conn, f) // handle connections concurrently
	}
	f.Close()
	return nil

}

func fileStream(c net.Conn, f *os.File) {
	fmt.Println("got to filestream")
	input := bufio.NewScanner(c)
	for input.Scan() {
		f.Write(input.Bytes())
		fmt.Printf("Received content %s\n", input.Text())
	}
}
