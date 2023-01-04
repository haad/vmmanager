package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DownloadHttpFile(url string) (string, error) {
	// Fetch the file from the web server
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching file:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the file contents
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading file contents:", err)
		return "", err
	}

	// Split the URL into its component parts
	parts := strings.Split(url, "/")
	path := parts[len(parts)-1]

	// Save the file to the local filesystem using the last component of the URL as the file name
	err = ioutil.WriteFile(path, contents, 0644)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return "", err
	}

	return path, nil
}
