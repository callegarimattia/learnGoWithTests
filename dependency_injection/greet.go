package greetings

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Received a request from ", r.RemoteAddr)
	Greet(w, "world")
}

func main() { //nolint:unused
	server := http.Server{ //nolint:exhaustruct
		Addr:        ":5001",
		Handler:     http.HandlerFunc(MyGreeterHandler),
		ReadTimeout: 5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
