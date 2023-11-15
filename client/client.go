package main

import (
	. "auction-replica-system/shared"
	"os"
)

var (
	// Resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func GetLeader() string {
	fileContents := GetFileContents(dockerId, "leader")

	if len(fileContents) > 0 {
		return fileContents[0]
	}

	return ""
}

func main() {
	_ = GetLeader()
}
