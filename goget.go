package goget

import (
	"bufio"
	"http"
	"io"
	"strings"
	"os"
)

const (
	bufSize         = 1024 * 8
	defaultFilename = "index.html"
)

func BuildAuthUrl(url string, username string, password string) (urlResult string, err os.Error) {
	urlObj, err := http.ParseURL(url)
	if err != nil {
		return "", err
	}
	if username == "" || password == "" {
		urlResult = url
		return
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

func GetOutfile(url string) (outfile string, err os.Error) {
	urlObj, err := http.ParseURL(url)
	if err != nil {
		return "", err
	}

	path := urlObj.Path
	slashIndex := strings.LastIndex(path, "/")
	if slashIndex == -1 {
		// no slash found, used a default filename
		outfile = defaultFilename
		return
	}
	slashIndex++
	outfile = path[slashIndex:len(path)]
	if outfile == "" {
		outfile = defaultFilename
	}
	return
}
