package main

import (
	"log"

	webserver "github.com/echodiv/simple_server/memory_list/internal/app/web_server"
)

var ADDRESS string = ":5000"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("<><> recovered in main function with exception: <><> ", r)
		}
	}()
	server := webserver.NewWebServer(ADDRESS)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
