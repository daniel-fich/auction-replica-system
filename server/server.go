package main

import (
	. "auction-replica-system/shared"
	"errors"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId      = os.Getenv("HOSTNAME")
	currentBid    = MakeThreadVariable[int64](0)
	currentLeader = ""
	isAuctionOver = MakeThreadVariable[bool](false)
)

func synchronize() {
	currentLeader, err := GetCurrentLeader()
	if err != nil {
		log.Fatal("Cannot start when there's no leader!")
	}

	for {
		// Only the leader will backup the other servers
		if dockerId == currentLeader {
			fileContents := GetFileContents(dockerId, NodesFilename)

			for _, backupHostname := range fileContents {
				// We don't want to backup to ourselves
				if backupHostname != dockerId {
					result, err := GetResult(backupHostname)

					if err != nil {
						log.Printf("Failed to get result on hostname %s\n", backupHostname)
						continue
					}

					currentHighest := currentBid.Get()

					if result.Outcome == currentHighest {
						continue
					}

					_, err = BidAmount(backupHostname, currentHighest)

					if err != nil {
						log.Printf("Failed to backup bid %d on hostname %s\n", currentHighest, backupHostname)
					}

					// TODO fail when X hostnames didn't backup
				}
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func randomCrash() {
	for {
		time.Sleep(5 * time.Second)
		shouldExit := rand.Int63n(10)

		if shouldExit == 1 {
			log.Fatal("Oh no! Something awful happened to my server, I spilled coffee all over it!")
		}
	}
}

func GetCurrentLeader() (string, error) {
	nodeIds := GetFileContents(dockerId, NodesFilename)
	if len(nodeIds) == 0 {
		return "", errors.New("No ids found")
	}
	return nodeIds[0], nil
}

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	WriteToSharedFile(dockerId, NodesFilename)

	go CheckPeers()
	go randomCrash()
	go synchronize()
	go startTimeServer(TimeService{})
	go PollTime()
	for {
	}
}
