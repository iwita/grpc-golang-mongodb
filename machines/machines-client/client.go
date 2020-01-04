package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-go-course/machines/machinespb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Machine Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := machinespb.NewMachineServiceClient(cc)

	//Create Machine
	machine := &machinespb.Machine{
		Processor:  "Intel i7 2.56 GHz",
		NumCores:   40,
		NumSockets: 2,
		Sockets: []*machinespb.Socket{
			&machinespb.Socket{
				NumCores:    20,
				L3CacheSize: 34000,
			},
			&machinespb.Socket{
				NumCores:    20,
				L3CacheSize: 34000,
			},
		},
		MemorySize: 128000,
	}
	createMachineRes, err := c.CreateMachine(context.Background(), &machinespb.CreateMachineRequest{Machine: machine})
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
	fmt.Printf("Machine has been created: %v", createMachineRes)
	machine_id := createMachineRes.GetMachine().GetMachineId()

	//Read Machine

	fmt.Println("Reading the machine")
	_, err2 := c.ReadMachine(context.Background(), &machinespb.ReadMachineRequest{
		MachineId: "5e0f4d2e90d3bcde47a05f83",
	})

	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readMachineReq := &machinespb.ReadMachineRequest{MachineId: machine_id}
	readMachineRes, readMachineErr := c.ReadMachine(context.Background(), readMachineReq)

	if readMachineErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readMachineErr)
	}

	fmt.Printf("Machine was read: %v \n", readMachineRes)

	//Update Machine

	fmt.Println("About to update a machine")
	newMachine := &machinespb.Machine{
		MachineId:  machine_id,
		Processor:  "Intel i7 2.56 GHz (updated)",
		NumCores:   41,
		NumSockets: 2,
		Sockets: []*machinespb.Socket{
			&machinespb.Socket{
				NumCores:    20,
				L3CacheSize: 34000,
			},
			&machinespb.Socket{
				NumCores:    20,
				L3CacheSize: 34000,
			},
		},
		MemorySize: 256000,
	}

	upRes, upErr := c.UpdateMachine(context.Background(), &machinespb.UpdateMachineRequest{Machine: newMachine})
	if upErr != nil {
		fmt.Printf("Error happened while updating :%v\n", upErr)
	}

	fmt.Printf("Machine was updated: %v", upRes)

	//Delete Machine

	fmt.Println("About to delete a machine")
	delRes, delErr := c.DeleteMachine(context.Background(), &machinespb.DeleteMachineRequest{MachineId: machine_id})

	if delErr != nil {
		fmt.Printf("Error happened while deleting: %v", delErr)
	}
	fmt.Printf("Machine was deleted: %v\n", delRes)

	//List Machines

	fmt.Println("About to List all the Machines")
	stream, err := c.ListMachine(context.Background(), &machinespb.ListMachineRequest{})
	if err != nil {
		log.Fatalf("error while calling ListMachine RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetMachine())
	}
}
