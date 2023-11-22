package main

import (
	. "auction-replica-system/shared"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type Remove interface {
	Remove(id string)
}

func CheckPeers() {
	for {
		contents := GetFileContents(dockerId, NodesFilename)

		for _, id := range contents {
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", id, strconv.Itoa(SERVER_PORT)), time.Millisecond*100)
			if err != nil {
				log.Println("Server is dead removing!:", id)
				// remove stuff
				continue
			}
			conn.Close()
		}
	}
}
