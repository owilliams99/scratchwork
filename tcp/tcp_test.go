package tcp

import (
	"io"
	"log"
	"net"
	"testing"
)

//var filecontent = "hello this worked correctly"

func TestTCPFileReceiver(t *testing.T) {
	//TODO Relative path needed for portability
	//filename := "/home/gspivey/Documents/Development/go/src/github.com/drayage/testdata/output/testfile.txt"
	filename := "/home/owilliams/go/src/github.com/owilliams99/scratchwork/test.txt"
	go TCPFileReceiver(filename)

	// Open connection and send file to our TCPFileReceiver
	conn, err := net.Dial("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	write := io.Writer(conn)
	write.Write([]byte(filecontent))
	conn.Close()

}
