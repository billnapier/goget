package main

import (
	"flag"
	"fmt"
	"goget"
	"os"
)

var username = flag.String("username", "", "Username for auth.")
var password = flag.String("password", "", "Password for auth.")

func main() {
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		url := flag.Arg(i)
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
	}
}
