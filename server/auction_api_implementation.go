package main

/*
This file is responsible for defining the remote procedure calls on the server
And also starting the GRPC server itself.
*/

import (
	proto "auction-replica-system/grpc"
	. "auction-replica-system/shared"
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedAuctionServiceServer
	id string
}

func (s *Server) Bid(ctx context.Context, in *proto.Amount) (*proto.Ack, error) {
	if currentBid.Get() < in.Amount {
		currentBid.Update(in.Amount)
	}
	synchronize()
	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}

func (s *Server) Result(ctx context.Context, in *proto.EmptyArgument) (*proto.Outcome, error) {
	if isAuctionOver.Get() {
		return &proto.Outcome{
			OutcomeType: proto.OutcomeType_RESULT,
			Outcome:     currentBid.Get(),
		}, nil
	} else {
		return &proto.Outcome{
			OutcomeType: proto.OutcomeType_HIGHEST_BID,
			Outcome:     currentBid.Get(),
		}, nil
	}
}

func startServer(server Server) {
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}

	log.Printf("Started server at port: %d\n", SERVER_PORT)

	// Register the grpc server and serve its listener
	proto.RegisterAuctionServiceServer(grpcServer, &server)
	serveError := grpcServer.Serve(listener)

	if serveError != nil {
		log.Fatal("Could not serve listener")
	}
}

func synchronize() {
	currentLeader, err := GetCurrentLeader()
	if err != nil {
		log.Fatal("Cannot start when there's no leader!")
	}

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
	}
}
