# PDB Downloader

This is a simple Go script that downloads PDB files from the RCSB PDB database. It prompts the user to enter one or more PDB IDs, downloads the corresponding PDB files, and saves them in a subdirectory named "pdb" in the current working directory.


## Installation

Make sure you have Go installed on your system. You can download it from the official website: https://golang.org/dl/

Create a new Go module for your project by running the following command in your terminal:


`go mod init github.com/<your-username>/<your-project-name>`

Install the pdb-go package by running the following command:

`go install github.com/shashankpritam/PracticingGo/pdb_go`


You can now use the package in your code to download PDB files. Don't forget to add error handling and proper error messages in case of any issues during the download process.

Note: The pdb-go package is currently available only on GitHub, so you need to have an active internet connection to download it.
## Usage

To use the PDB Downloader, simply run the following command:

`./main.go`

The script will prompt you to enter one or more PDB IDs, separated by spaces. For example:



`Enter PDB IDs separated by spaces:
1ake 1brs 1dug` 

The script will then download the corresponding PDB files and save them in a subdirectory named "pdb" in the current working directory.

## Credits

This script was created by Shashank Pritam with the help of ChatGPT, a language model trained by OpenAI. It is based on the [RCSB PDB RESTful web service](https://www.rcsb.org/pages/webservices/restful) and the [Go standard library](https://golang.org/pkg/).