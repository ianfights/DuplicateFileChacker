package main
// This program was originally created by myself in early 2022 as a utility tool, however I have adapted it to fit the requirments set forth by the College Board for this class
// The repo can be found here https://github.com/ianfights/DuplicateFileChacker
// I hope this alleviates any confusion from the inevitable result of you finding my GitHub profile
// Code from built-in Go libraries
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
	checkDupFiles(fileHashes, files)

	// Check if the file hash is already in the array and if so delete the file

	// We don't need to loop through files as their position in the fileHash array will be the same as their pos in files

}

func checkDupFiles(fileHashes[]string, files[]string){
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

		fmt.Printf("You will be deleting %v files. Is this correct?\n Y or N\n", len(duplicateFiles)/2)
		var delCheck string
		fmt.Scan(&delCheck)
		if delCheck == "N" || delCheck == "n" {
			fmt.Println("Exiting")
			return
		}

		// Delete files
		for i := range duplicateFiles {
			// Mod by two so only a single dup file is deleted
			// is there a better way to do this? Undoubtably, but I really could care less because this works
			if i%2 == 0 {
				e := os.Remove(duplicateFiles[i])
				if e != nil {
					log.Fatal(e)
				}
				fmt.Printf("Removed %v", duplicateFiles[i])
			}
		}
	} else {
		fmt.Println("No duplicate files")
		return
	}
}
// This method taken from https://medium.com/@manigandand/list-files-in-a-directory-using-golang-963b1df11304
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

// Portions of this method taken from https://stackoverflow.com/questions/15879136/how-to-calculate-sha256-file-checksum-in-go
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
