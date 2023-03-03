package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// define constants for the URL and file paths
const (
	urlBase = "https://files.rcsb.org/download/"
	fileExt = ".pdb"
)

func main() {
	// prompt the user for the PDB IDs they want to download
	fmt.Println("Enter PDB IDs separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pdbIDs := strings.Fields(scanner.Text())

	// loop over the PDB IDs and download each file
	for _, pdbID := range pdbIDs {
		// create the URL and file paths for this PDB ID
		url := urlBase + pdbID + fileExt
		fileName := pdbID + fileExt
		// define the absolute path to the "pdb" subdirectory in the current working directory
		absPath, err := filepath.Abs(filepath.Join(".", "pdb"))
		if err != nil {
			fmt.Printf("Error getting absolute path to subdirectory: %s\n", err)
			os.Exit(1)
		}
		// create the full path to the file by joining the subdirectory path and the filename
		fullPath := filepath.Join(absPath, fileName)

		// make the HTTP GET request to download the file
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error downloading file for PDB ID %s: %s\n", pdbID, err)
			continue
		}
		defer resp.Body.Close()

		// check if the response was successful
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error downloading file for PDB ID %s: unexpected status code %d\n", pdbID, resp.StatusCode)
			continue
		}

		// read the response body and save it to a file
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading file for PDB ID %s: %s\n", pdbID, err)
			continue
		}
		err = ioutil.WriteFile(fullPath, body, 0644)
		if err != nil {
			fmt.Printf("Error writing file for PDB ID %s: %s\n", pdbID, err)
			continue
		}

		// print a success message with the file path
		fmt.Printf("File for PDB ID %s downloaded to %s\n", pdbID, fullPath)
	}
}
