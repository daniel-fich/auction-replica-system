package main

import (
	. "auction-replica-system/shared"
	"os"
)

var (
	// HOSTNAME resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	WriteToSharedFile(dockerId, NodesFilename)
}
