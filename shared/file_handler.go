package shared

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	NodesFilename  = "nodes"
	LeaderFilename = "leader"
)

func WriteToSharedFile(dockerId string, filename string) {
	sharedFilePath := GetPath(filename)

	f, err := os.OpenFile(sharedFilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		log.Println("[WriteToSharedFile] Couldn't open the shared file! On server ", dockerId)
	}
	defer f.Close()
	if _, err := f.WriteString(dockerId + "\n"); err != nil {
		log.Println("Couldn't write to shared file! On server ", dockerId)
	}
}

func GetFileContents(dockerId string, filename string) []string {
	sharedFilePath := GetPath(filename)

	if _, err := os.Stat(sharedFilePath); os.IsNotExist(err) {
		log.Println("File does not exist: " + sharedFilePath)
	}

	f, err := os.OpenFile(sharedFilePath,
		os.O_RDONLY,
		0644)

	if err != nil {
		log.Println("[GetFileContents] Couldn't open the shared file! On server ", dockerId)
	}
	defer f.Close()
	bytes := make([]byte, 1_000_000)

	f.Read(bytes)
	contents := string(bytes)
	strs := strings.Split(contents, "\n")

	if len(strs) > 0 {
		// Removes whitespace from end of file
		strs = strs[:len(strs)-1]
	}
	return strs
}

func GetPath(file string) string {
	return filepath.FromSlash("/var/hosts/" + file)
}
