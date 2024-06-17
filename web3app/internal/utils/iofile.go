package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectRoot := cwd
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, ".git")); os.IsNotExist(err) {
			projectRoot = filepath.Dir(projectRoot)
			if projectRoot == "/" || projectRoot == "." {
				return "", fmt.Errorf("project root not found")
			}
		} else {
			break
		}
	}

	return projectRoot, nil
}

func WriteStringToFile(input string, fname string) bool {
	file, errs := os.Create(fname)
	if errs != nil {
		log.Println("Failed to create file:", errs)
		return false
	}
	defer file.Close()

	_, errs = file.WriteString(input)
	if errs != nil {
		log.Println("Failed to write to file:", errs) //print the failed message
		return false
	}
	fmt.Println("Wrote to file " + fname) //print the success message
	return true
}
