package main

import (
	"bytes"
	"flag"
	"fmt"
	"http"
	"io"
	"os"
)

func ReaderToString(reader io.Reader) string {
	buffer := bytes.NewBufferString("")
	buffer.ReadFrom(reader)
	return buffer.String()
}

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

func FetchUrl(url string) (result string, err os.Error) {
	r, _, err := http.Get(url)
	if err != nil {
		return "", err
        }
	result = ReaderToString(r.Body)
	return 
}

func WriteOutData(data string, outfile string) (err os.Error) {
	file, err := os.Open(outfile, os.O_WRONLY | os.O_CREAT, 0777)
	if err != nil {
 		return err
	}
	file.WriteString(data)
	file.Close()
	return 
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
		fmt.Println("Error building auth url");
		os.Exit(1)
	}
	data, err := FetchUrl(authUrl)
	if err != nil {
		fmt.Println("Error fetching url");
		os.Exit(1)
	}	
	err = WriteOutData(data, *outfile)
	if err != nil {
		fmt.Println("Error writing out data: " + err.String());
		os.Exit(1)
	}
}
