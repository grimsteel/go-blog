package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	listen_address := os.Getenv("LISTEN_ADDRESS")
	if len(listen_address) == 0 {
		// set a default address
		listen_address = "127.0.0.1:8080"
	}

  mux := http.NewServeMux()

	// serve static files
	staticPath := "/static/"
  mux.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(http.Dir("./static"))))
	
  mux.HandleFunc("/", index)

  server := &http.Server{
    Addr:     listen_address,
    Handler:  mux,
  }

	// start the server
	fmt.Printf("Listening on %s\n", listen_address)
  server.ListenAndServe()
}
