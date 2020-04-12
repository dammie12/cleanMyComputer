package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	homeDir := getUserDirectory()
	baseDir := getRootDirectory(homeDir)
	changeDirectory(baseDir)
	allFiles, _ := readDir(baseDir)
	fmt.Println(allFiles[0])
}

func getUserDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error retrieving home directory", err)
		os.Exit(1)
	}
	if homeDir == "" {
		fmt.Println("Home directory cannot be empty", homeDir)
		os.Exit(1)
	}
	fmt.Println("Home directory is", homeDir)
	return homeDir
}

func getRootDirectory(s string) string {
	var home string
	f := filepath.ToSlash(s)
	strSlice := strings.Split(f, "/")
	if len(strSlice) > 0 {
		home = strSlice[0]
	}
	fmt.Println("Root directory is", home)
	return home
}

func changeDirectory(s string) {
	_, err := os.Stat(s)
	if err != nil {
		fmt.Println("Error changing to this directory:- ", s)
		os.Exit(1)
	}
	os.Chdir(s)
	fmt.Println("Directory changed to", s)
}

func readDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

/*
func writeToFile(root string) ([]string, error) {
	err := ioutil.WriteFile("temp.txt", []byte(linesToWrite), 0777)
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
*/
