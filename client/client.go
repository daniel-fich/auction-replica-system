package main

import (
	. "auction-replica-system/shared"
	"fmt"
	"os"
)

var (
	// HOSTNAME resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func GetLeader() string {
	fileContents := GetFileContents(dockerId, NodesFilename)

	if len(fileContents) > 0 {
		return fileContents[0]
	}

	return ""
}

func main() {
	test := GetLeader()

	fmt.Println(test)
}
