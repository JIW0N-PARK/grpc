package main

import(
	"net"
	"log"
	"context"
	"time"
	"fmt"
	// "path/filepath"
	"io/ioutil"
	"os"

	"google.golang.org/grpc"
	pb "../proto/test"

	"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"gopkg.in/yaml.v2"

)

type server struct {
	pb.UnimplementedWebServer
}

type Post struct {
	Name string `json:"name,omitempty"`
	Age int32 `json:"body,omitempty"`
}

type DatabaseConfig struct {
	dbName string `yaml:"dbName"`
	dbType string `yaml:"dbType"`
	user string `yaml:"user"`
	password string `yaml:"password"`
}

const(
	port = ":50051"
)

func (s server) Register(ctx context.Context, in *pb.RegisterReq) (*pb.Response, error) {
	log.Printf("Register")
	log.Printf("Name=%s\n", in.GetName())
	log.Printf("Age=%d\n", in.GetAge())
	InsertPost(in.GetName(), in.GetAge())
	return &pb.Response{Message: "Complete!"}, nil
}

func (s server) Check(ctx context.Context, in *pb.CheckReq) (*pb.Response, error){
	res := GetPost(in.GetName())
	log.Printf("Retrive")
	log.Printf("Name=%s\n", in.GetName())
	log.Printf("Result - %s\n", res)
	return &pb.Response{Message: res}, nil
}

func InsertPost(name string, age int32) {
	client := ConnectDB()

	post := Post{name, age}

	collection := client.Database("test").Collection("grpc")
	
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted post with ID:", insertResult.InsertedID)	
}

func GetPost(name string) (string) {
	client := ConnectDB()

	collection := client.Database("test").Collection("grpc")
	
	filter := bson.D{{"name", name}}
	
	var post Post

	err := collection.FindOne(context.TODO(), filter).Decode(&post)
	
	if err != nil {
		return "Not Found"
	} else {return "Found"}	
}

func ConnectDB()(client *mongo.Client){

	// filename, _ := filepath.Abs("../databaseConfig.yml")
	// yamlFile, err := ioutil.ReadFile(filename)
	
	var databaseConfig DatabaseConfig
	reader, _ := os.Open("/../databaseConfig.yml")
  buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &databaseConfig)

	// yamlErr := yaml.Unmarshal(yamlFile, databaseConfig)

	// if yamlErr != nil {
	// 	panic(yamlErr)
	// }

	fmt.Printf("%+v\n", databaseConfig)

	client, clientErr := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://"+databaseConfig.user+":1029@cluster0.rijup.gcp.mongodb.net/test?retryWrites=true&w=majority"))
	
	if clientErr != nil {
			log.Fatal(clientErr)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
					log.Fatal(err)
	}
	return client
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWebServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}