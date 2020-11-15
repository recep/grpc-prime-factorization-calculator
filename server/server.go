package main

import (
	"log"
	"net"

	"github.com/recep/grpc-prime-factorization-calculator/pb"
	"google.golang.org/grpc"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

func main() {

	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFactorizationServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}

type server struct {
	pb.UnimplementedFactorizationServiceServer
}

func (s *server) Separate(req *pb.NumberRequest, stream pb.FactorizationService_SeparateServer) error {
	num := req.GetNumber()

	for k := 2; 1 < num; {
		if num%int32(k) == 0 {
			res := &pb.PrimeNumResponse{PrimeNumber: int32(k)}
			num = num / int32(k)
			err := stream.Send(res)
			if err != nil {
				return err
			}
			continue
		}
		k++
	}
	return nil
}
