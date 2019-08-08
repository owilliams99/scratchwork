package http

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"net/url"
)

func HTTPFileReceiver(file string) (err error) {

	/*handler := func(w http.ResponseWriter, r *http.Request) {
		//fileStream()
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	*/
	//f, _ := os.Create(file)
	//defer f.Close()

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	go log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//http.ServeContent(rw, r, file, time.Now(), bytes.NewReader(data))
		params, err := url.ParseQuery(r.URL.Path)

	})))

	//go fetch("http://localhost:8000")

	return nil
}

func fileStream(c net.Conn, f *os.File) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		f.Write(input.Bytes())
		fmt.Printf("Received content %s\n", input.Text())
	}
}

func fetch(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)

}
