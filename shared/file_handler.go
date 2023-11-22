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
	log.Printf("The server ip address is: %s\n", dockerId)

	f, err := os.OpenFile(filepath.Join("var", "hosts", filename),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		log.Println("Couldn't open the shared file! On server ", dockerId)
	}
	defer f.Close()
	if _, err := f.WriteString(dockerId + "\n"); err != nil {
		log.Println("Couldn't write to shared file! On server ", dockerId)
	}
}

func GetFileContents(dockerId string, filename string) []string {
	f, err := os.OpenFile(filepath.Join("var", "hosts", filename),
		os.O_RDONLY,
		0644)

	if err != nil {
		log.Println("Couldn't open the shared file! On server ", dockerId)
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
