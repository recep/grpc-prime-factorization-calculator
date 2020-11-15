package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/recep/grpc-prime-factorization-calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewFactorizationServiceClient(conn)

	fmt.Printf("Enter number here: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil || num < 2 {
			fmt.Printf("Please enter number: ")
			continue
		}

		req := &pb.NumberRequest{
			Number: int32(num),
		}

		stream, err := client.Separate(context.Background(), req)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", strings.Repeat("-", 50))
		i := 1
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("%6d |\n", 1)
				break
			}

			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%6d | %2d\n", num/i, res.PrimeNumber)
			i *= int(res.PrimeNumber)
		}
		fmt.Printf("Enter number here: ")
	}
}
