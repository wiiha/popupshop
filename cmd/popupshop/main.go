package main

import (
	"flag"
	"log"

	"github.com/wiiha/popupshop/files"
	"github.com/wiiha/popupshop/server"
)

func main() {

	/*
		Read flags from command line
	*/
	help := flag.Bool("help", false, "Show help message")
	addr := flag.String("addr", "localhost", "Address to listen on")
	port := flag.String("port", "8080", "Port to listen on")
	dirToServe := flag.String("root", ".", "Root directory for files")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	fileService := files.NewFileService(*dirToServe)

	fAPI := server.NewFilesAPI(*addr, *port, fileService)
	err := fAPI.Run()
	if err != nil {
		log.Fatalf("when running server: %v", err)
	}

}
