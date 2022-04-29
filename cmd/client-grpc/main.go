package main

import (
	"context"
	"flag"
	v1 "github.com/karimbayli/mgid-task/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewEmployeeServiceClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Call GetSortedEmployees sorted by birth year
	req1 := v1.GetSortedEmployeesRequest{Field: "yearOfBirth"}
	res1, err := c.GetSortedEmployees(ctx1, &req1)
	if err != nil {
		log.Printf("GetSortedEmployees sorted by birth year failed: %v", err)
	}
	log.Printf("GetSortedEmployees sorted by birth year result: <%+v>\n\n", res1)

	// Call GetSortedEmployees sorted by name
	ctx2, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req2 := v1.GetSortedEmployeesRequest{Field: "name"}
	res2, err := c.GetSortedEmployees(ctx2, &req2)
	if err != nil {
		log.Printf("GetSortedEmployees sorted by name failed: %v", err)
	}
	log.Printf("GetSortedEmployees sorted by name result: <%+v>\n\n", res2)

	// Call GetSortedEmployees sorted by salary
	ctx3, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req3 := v1.GetSortedEmployeesRequest{Field: "salary"}
	res3, err := c.GetSortedEmployees(ctx3, &req3)
	if err != nil {
		log.Printf("GetSortedEmployees sorted by salary failed: %v", err)
	}
	log.Printf("GetSortedEmployees sorted by salary result: <%+v>\n\n", res3)

	// Call GetImitatedTimeout
	ctx4, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	defer cancel()

	req4 := v1.GetImitatedTimeoutRequest{}
	res4, err := c.GetImitatedTimeout(ctx4, &req4)
	if err != nil {
		log.Printf("GetImitatedTimeout failed: %v", err)
	}
	log.Printf("GetImitatedTimeout result: <%+v>\n\n", res4)

	// Call GetMedianOfSalaries
	ctx5, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	req5 := v1.GetMedianOfSalariesRequest{}
	res5, err := c.GetMedianOfSalaries(ctx5, &req5)
	if err != nil {
		log.Printf("GetMedianOfSalaries failed: %v", err)
	}
	log.Printf("GetMedianOfSalaries result: <%+v>\n\n", res5)

	// Call GetAverageOfSalariesRequest
	ctx6, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req6 := v1.GetAverageOfSalariesRequest{}
	res6, err := c.GetAverageOfSalaries(ctx6, &req6)
	if err != nil {
		log.Printf("GetAverageOfSalaries failed: %v", err)
	}
	log.Printf("GetMedianOfSalaries result: <%+v>\n\n", res6)

	// Call GetEldestEmployee
	ctx7, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req7 := v1.GetEldestEmployeeRequest{}
	res7, err := c.GetEldestEmployee(ctx7, &req7)
	if err != nil {
		log.Printf("GetEldestEmployee failed: %v", err)
	}
	log.Printf("GetEldestEmployee result: <%+v>\n\n", res7)

	// Call GetHighestPaid
	ctx8, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req8 := v1.GetHighestPaidRequest{}
	res8, err := c.GetHighestPaid(ctx8, &req8)
	if err != nil {
		log.Printf("GetHighestPaid failed: %v", err)
	}
	log.Printf("GetHighestPaid result: <%+v>\n\n", res8)
}
