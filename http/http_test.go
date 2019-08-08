package http

import (
	"testing"
)

var filecontent = `---
receipt:     Oz-Ware Purchase Invoice
date:        2012-08-06
customer:
    first_name:   Dorothy
    family_name:  Gale

items:
    - part_no:   A4786
      descrip:   Water Bucket (Filled)
      price:     1.47
      quantity:  4

    - part_no:   E1628
      descrip:   High Heeled "Ruby" Slippers
      size:      8
      price:     133.7
      quantity:  1

bill-to:  &id001
    street: |
            123 Tornado Alley
            Suite 16
    city:   East Centerville:w

    state:  KS

ship-to:  *id001

specialDelivery:  >
    Follow the Yellow Brick
    Road to the Emerald City.
    Pay no attention to the
    man behind the curtain.
...`

func TestHTTPFileReceiver(t *testing.T) {
	filename := "/home/owilliams/go/src/github.com/owilliams99/scratchwork/test.txt"

	//go
	HTTPFileReceiver(filename)

	/*conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	write := io.Writer(conn)
	write.Write([]byte(filecontent))
	conn.Close()
	*/
}
