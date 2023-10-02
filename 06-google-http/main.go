package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Get the status
	fmt.Println(resp.Status)

	// Read the body (manually)
	// body := make([]byte, 99999)
	// bytes, _ := resp.Body.Read(body)
	// if bytes != 0 {
	// 	fmt.Println(string(body))
	// }

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes: ", len(bs))
	return len(bs), nil
}
