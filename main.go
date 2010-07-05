package main

import (
	"flag"
	"fmt"
	"goget"
	"os"
)

var username = flag.String("username", "", "Username for auth.")
var password = flag.String("password", "", "Password for auth.")

func GoGet(url string) chan string {
	ch := make(chan string)
	go func() {
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
		ch <- outfile
	}()
	return ch
}

func main() {
	flag.Parse()

	results := make([]chan string, flag.NArg())

	// Kick off all the async downloads
	for i := 0; i < flag.NArg(); i++ {
		results[i] = GoGet(flag.Arg(i))
	}

	// And wait for their completion
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println("Downloaded to " + <- results[i])
	}
}
