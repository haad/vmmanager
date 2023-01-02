package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var url string
	var fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch a qcow2 disk image from a web server",
		Long:  "Fetch a qcow2 disk image from a web server",
		Run: func(cmd *cobra.Command, args []string) {
			// Fetch the file from the web server
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error fetching file:", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			// Read the file contents
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading file contents:", err)
				os.Exit(1)
			}

			// Split the URL into its component parts
			parts := strings.Split(url, "/")

			// Save the file to the local filesystem using the last component of the URL as the file name
			err = ioutil.WriteFile(parts[len(parts)-1], contents, 0644)
			if err != nil {
				fmt.Println("Error saving file:", err)
				os.Exit(1)
			}

			fmt.Println("File saved successfully!")
		},
	}

	fetchCmd.Flags().StringVarP(&url, "url", "U", "", "URL to fetch qcow2 image from.")
	fetchCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(fetchCmd)
}
