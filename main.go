package main

import (
	"flag"
	"fmt"
	"goget"
	"os"
)

var username = flag.String("username", "", "Username for auth.")
var password = flag.String("password", "", "Password for auth.")

func GoGet(url string, resultChan chan string) {
	outfile, err := goget.GetOutfile(url)
	if err != nil {
		fmt.Println("Error determining outfile: " + err.String())
		os.Exit(1)
	}

	authUrl, err := goget.BuildAuthUrl(url, *username, *password)
	if err != nil {
		fmt.Println("Error building auth url: " + err.String())
		os.Exit(1)
	}
	err = goget.FetchUrl(authUrl, outfile)
	if err != nil {
		fmt.Println("Error fetching url: " + err.String())
		os.Exit(1)
	}
	resultChan <- outfile
}

func main() {
	flag.Parse()

	results := make(chan string)

	// Kick off all the async downloads
	for i := 0; i < flag.NArg(); i++ {
		go GoGet(flag.Arg(i), results)
	}

	// And wait for their completion
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println("Downloaded to " + <-results)
	}
}
