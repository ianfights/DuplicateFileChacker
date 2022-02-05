package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Please enter the path you want to search")
	var path string = "./"
	fmt.Scan(&path)

	fmt.Printf("You will be searching %v for duplicte files. \n Is this corrent? \n Y or N\n", path)
	var check string
	fmt.Scan(&check)
	if check == "N" || check == "n" {
		fmt.Println("Please rerun the program again with the correct file path")
		return
	}

	files, err := scanDirs(path)

	if err != nil {
		fmt.Println(err)
		return
	}
	var fileHashes []string
	for _, file := range files {
		fileHashes = append(fileHashes, getFileHash(file))
	}
	// for i, hash := range fileHashes {
	// 	fmt.Println("File Path:", files[i], "File Hash:", hash)
	// }
	// Check if the file hash is already in the array and if so delete the file

	// We don't need to loop through files as their position in the fileHash array will be the same as their pos in files

	var duplicateFiles []string

	m := make(map[string]int)

	for i, iFileHash := range fileHashes {
		for _, fileHash := range fileHashes {
			if iFileHash == fileHash {
				m[files[i]]++
				if m[files[i]] == 2 {
					// Add it to the dupFiles slice
					duplicateFiles = append(duplicateFiles, files[i])
					fmt.Println("Found a duplicate file")
					fmt.Println("File Path:", files[i], "File Hash:", iFileHash)
				}
			}
		}
	}
	if len(duplicateFiles) != 0 {

		fmt.Printf("You will be deleting %v files. Is this correct?\n Y or N\n", len(duplicateFiles))
		var delCheck string
		fmt.Scan(&delCheck)
		if check == "N" || check == "n" {
			fmt.Println("Exiting")
			return
		}

		// Delete files
		for _, file := range duplicateFiles {
			e := os.Remove(file)
			if e != nil {
				log.Fatal(e)
			}
			fmt.Printf("Removed %v", file)
		}
	} else {
		fmt.Println("No duplicate files")
		return
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func scanDirs(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getFileHash(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}
