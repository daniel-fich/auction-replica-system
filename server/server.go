package main

import (
	. "auction-replica-system/shared"
	"log"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId   = os.Getenv("HOSTNAME")
	currentBid = MakeThreadVariable[int64](0)
)

func synchronize() {
	for {
		fileContents := GetFileContents(dockerId, NodesFilename)

		for _, backupHostname := range fileContents {
			if backupHostname != dockerId {
				// ping	id here

				result, err := getResult(backupHostname)

				if err != nil {
					log.Printf("Failed to get result on hostname %s\n", backupHostname)
					continue
				}

				currentHighest := currentBid.Get()

				if result == currentHighest {
					log.Printf("Bid matches between %s and %s, bid: %d\n", dockerId, backupHostname, currentHighest)
					continue
				}

				_, err = bidAmount(backupHostname, currentHighest)

				if err != nil {
					log.Printf("Failed to backup bid %d on hostname %s\n", currentHighest, backupHostname)
				}

				// TODO fail when X hostnames didn't backup
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	time.Sleep(time.Second)
	WriteToSharedFile(dockerId, NodesFilename)

	go synchronize()
	// go CheckPeers()

	for {

	}
}
