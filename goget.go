package main

import (
	"bufio"
	"flag"
	"fmt"
	"http"
	"io"
	"os"
)

const (
	bufSize = 1024 * 8
)

func BuildAuthUrl(url string, username string, password string) (urlResult string, err os.Error) {
	urlObj, err := http.ParseURL(url)
	if err != nil {
		return "", err
	}
	// Fill in the userinfo field
	urlObj.Userinfo = username + ":" + password
	urlResult = urlObj.String()
	return
}

func FetchUrl(url string, outfile string) (err os.Error) {
	r, _, err := http.Get(url)
	if err != nil {
		return err
	}
	file, err := os.Open(outfile, os.O_WRONLY|os.O_CREAT, 0777)
	if err != nil {
		return err
	}
	bufferedWriter, err := bufio.NewWriterSize(file, bufSize)
	if err != nil {
		return err
	}

	io.Copy(bufferedWriter, r.Body)
	r.Body.Close()
	file.Close()

	return nil
}

var username = flag.String("username", "", "Username for auth.")
var password = flag.String("password", "", "Password for auth.")
var url = flag.String("url", "", "URL to fetch")
var outfile = flag.String("o", "", "Where to save the file")

func main() {
	flag.Parse()

	if *url == "" {
		fmt.Println("Must specify url to fetch")
		os.Exit(1)
	}

	authUrl, err := BuildAuthUrl(*url, *username, *password)
	if err != nil {
		fmt.Println("Error building auth url: " + err.String())
		os.Exit(1)
	}
	err = FetchUrl(authUrl, *outfile)
	if err != nil {
		fmt.Println("Error fetching url: " + err.String())
		os.Exit(1)
	}
}
