package main

import (
	. "auction-replica-system/shared"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId   = os.Getenv("HOSTNAME")
	currentBid = MakeThreadVariable[int64](0)
)

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	time.Sleep(time.Second)
	WriteToSharedFile(dockerId, NodesFilename)
	// go CheckPeers()

	for {

	}
}
