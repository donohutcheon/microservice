//https://www.youtube.com/watch?v=wxkEQxvxs3w&list=PLijOg6JCGTRbmg_LOqIRs5vwQP2c3fMKg

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/donohutcheon/microservice/homepage"
	"github.com/donohutcheon/microservice/server"
)

var (
	//CertFile environment variable for CertFile
	CertFile = os.Getenv("MICROSERVICE_CERT_FILE")
	//KeyFile environment variable for KeyFile
	KeyFile = os.Getenv("MICROSERVICE_KEY_FILE")
	//ServiceAddress address to listen on
	ServiceAddress = os.Getenv("MICROSERVICE_SERVICE_ADDRESS")
)

func main() {
	logger := log.New(os.Stdout, "microservice", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	//https://blog.cloudflare.com/exposing-go-on-the-internet/

	srv := server.New(mux, ServiceAddress)
	logger.Println("Starting the server...")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
		fmt.Println(err)
	}
	fmt.Println("Done!")
}
