package fetch_pdb

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func getStructure(inputPdbGiven string) (*Structure, error) {
	// First, try to open the PDB file
	inputPdb := filepath.Join("database_folder", inputPdbGiven+".pdb")
	inputStructure, err := parser.GetStructureFromFile(inputPdbGiven, inputPdb)
	if err == nil {
		return inputStructure, nil
	}
	
	// If the file is not found or there is an error, try to download it
	inputPdbUrl := "https://files.rcsb.org/download/" + inputPdbGiven + ".pdb"
	response, err := http.Get(inputPdbUrl)
	if err != nil {
		return nil, errors.New("failed to download PDB file")
	}
	defer response.Body.Close()
	
	// Save the downloaded PDB file to the database_folder directory
	inputPdb = filepath.Join("database_folder", inputPdbGiven+".pdb")
	if err := ioutil.WriteFile(inputPdb, response.Body, 0644); err != nil {
		return nil, errors.New("failed to save PDB file")
	}
	
	// Parse the downloaded PDB file
	inputStructure, err = parser.GetStructureFromFile(inputPdbGiven, inputPdb)
	if err != nil {
		return nil, errors.New("failed to parse PDB file")
	}
	return inputStructure, nil
}
