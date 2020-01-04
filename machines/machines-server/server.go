package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc/reflection"

	"go.mongodb.org/mongo-driver/bson"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/grpc-go-course/machines/machinespb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

func (*server) ListMachine(req *machinespb.ListMachineRequest, stream machinespb.MachineService_ListMachineServer) error {

	fmt.Println("List Machines")

	// Don't need to provide a filter, we return all the machines
	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal Error: %v", err),
		)
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &MachineItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Decoding error: %v", err),
			)
		}
		stream.Send(&machinespb.ListMachineResponse{Machine: dataToMachinePb(data)})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal Error: %v", err),
		)
	}
	return nil
}

func (*server) DeleteMachine(ctx context.Context, req *machinespb.DeleteMachineRequest) (*machinespb.DeleteMachineResponse, error) {
	fmt.Println("Delete Machine")

	oid, err := primitive.ObjectIDFromHex(req.GetMachineId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprint("Cannot parse ID"),
		)
	}
	filter := bson.M{"_id": oid}
	DelRes, DelErr := collection.DeleteOne(context.Background(), filter)

	if DelErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in mongodb: %v", DelErr),
		)
	}
	if DelRes.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find machine in mongoDB: %v", DelErr),
		)
	}

	return &machinespb.DeleteMachineResponse{
		MachineId: req.GetMachineId(),
	}, nil

}

func (*server) UpdateMachine(ctx context.Context, req *machinespb.UpdateMachineRequest) (*machinespb.UpdateMachineResponse, error) {

	fmt.Println("Update Machine")
	machine := req.GetMachine()
	oid, err := primitive.ObjectIDFromHex(machine.GetMachineId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprint("Cannot parse ID"),
		)
	}

	//Create an empty struct
	data := &MachineItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find machine with specified ID: %v", err),
		)
	}

	//Update our internal struct
	data.Processor = machine.GetProcessor()
	data.Num_cores = machine.GetNumCores()
	data.Num_sockets = machine.GetNumSockets()

	_, Uperr := collection.ReplaceOne(context.Background(), filter, data)
	if Uperr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in mongodb: %v", Uperr),
		)
	}

	return &machinespb.UpdateMachineResponse{
		Machine: dataToMachinePb(data),
	}, nil
}

func dataToMachinePb(data *MachineItem) *machinespb.Machine {

	var sockets []*machinespb.Socket
	for _, sock := range data.Sockets {
		s := &machinespb.Socket{
			NumCores:    sock.Num_cores,
			L3CacheSize: sock.L3_cache_size,
		}
		sockets = append(sockets, s)
	}

	return &machinespb.Machine{
		MachineId:  data.ID.Hex(),
		Processor:  data.Processor,
		NumCores:   data.Num_cores,
		NumSockets: data.Num_sockets,
		Sockets:    sockets,
	}
}

func (*server) ReadMachine(ctx context.Context, req *machinespb.ReadMachineRequest) (*machinespb.ReadMachineResponse, error) {

	fmt.Println("Read Machine Request")

	machineID := req.GetMachineId()
	oid, err := primitive.ObjectIDFromHex(machineID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprint("Cannot parse ID"),
		)
	}
	//Create an empty struct
	data := &MachineItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find machine with specified ID: %v", err),
		)
	}

	var sockets []*machinespb.Socket
	for _, sock := range data.Sockets {
		s := &machinespb.Socket{
			NumCores:    sock.Num_cores,
			L3CacheSize: sock.L3_cache_size,
		}
		sockets = append(sockets, s)
	}
	//data.Sockets = sockets

	return &machinespb.ReadMachineResponse{
		Machine: dataToMachinePb(data),
	}, nil

}

func (*server) CreateMachine(ctx context.Context, req *machinespb.CreateMachineRequest) (*machinespb.CreateMachineResponse, error) {

	fmt.Println("Creare Machine Request")
	machine := req.GetMachine()
	sockets := machine.GetSockets()
	var s []SocketItem
	for _, socket := range sockets {
		socket := SocketItem{
			Num_cores:     socket.GetNumCores(),
			L3_cache_size: socket.GetL3CacheSize(),
		}
		s = append(s, socket)
	}
	data := MachineItem{
		Processor:   machine.GetProcessor(),
		Num_cores:   machine.GetNumCores(),
		Num_sockets: machine.GetNumSockets(),
		Sockets:     s,
		Memory_size: machine.GetMemorySize(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v ", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Cannot convert to OID: %v ", err),
		)
	}

	return &machinespb.CreateMachineResponse{
		Machine: &machinespb.Machine{
			MachineId:  oid.Hex(),
			Processor:  machine.GetProcessor(),
			NumCores:   machine.GetNumCores(),
			NumSockets: machine.GetNumSockets(),
			Sockets:    machine.GetSockets(),
			MemorySize: machine.GetMemorySize(),
		},
	}, nil
}

type SocketItem struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Num_cores     int32              `bson:"num_cores"`
	L3_cache_size int64              `bson:"l3_cache_size"`
}

type MachineItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Processor   string             `bson:"pr ocessor"`
	Num_cores   int32              `bson:"num_cores"`
	Num_sockets int32              `bson:"num_sockets"`
	Sockets     []SocketItem       `bson:"sockets"`
	Memory_size int64              `bson:"mem_size"`
}

func main() {
	//if we crash the go code, we get the file name and row number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Connect to mongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//Use a Collection

	collection = client.Database("mydb").Collection("machines")

	fmt.Println("Machine Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	machinespb.RegisterMachineServiceServer(s, &server{})

	//Implement Reflection in order to use EVANS CLI
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//Wait for Contl-C to exit
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	//Block until a signal is received
	<-ch

	fmt.Println("Stoping the Server")
	s.Stop()
	fmt.Println("Closing the Listener")
	lis.Close()
	fmt.Println("Closing mongoDB Connection")
	client.Disconnect(context.TODO())
}
