package main

import(
	"net"
	"log"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"google.golang.org/grpc"
	pb "../proto/test"

	_ "go.mongodb.org/mongo-driver/bson"
  _ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	
	"gopkg.in/yaml.v2"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type server struct {
	pb.UnimplementedWebServer
}

type Post struct {
	Name string `json:"name,omitempty"`
	Age int32 `json:"body,omitempty"`
}

type Config struct {
	DBName string `yaml:"dbName"`
	DBType string `yaml:"dbType"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Port string `yaml:"port"`
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

// func (s server) Check(ctx context.Context, in *pb.CheckReq) (*pb.Response, error){
// 	res := GetPost(in.GetName())
// 	log.Printf("Retrive")
// 	log.Printf("Name=%s\n", in.GetName())
// 	log.Printf("Result - %s\n", res)
// 	return &pb.Response{Message: res}, nil
// }

func InsertPost(name string, age int32) {
	dbClient := ConnectDB()

	result, err := dbClient.Exec("INSERT INTO grpc VALUES (?, ?, ?)", 4, name, age)	

	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}

// func GetPost(name string) (string) {
// 	client := ConnectDB()

// 	collection := client.Database("test").Collection("grpc")
	
// 	filter := bson.D{{"name", name}}
	
// 	var post Post

// 	err := collection.FindOne(context.TODO(), filter).Decode(&post)
	
// 	if err != nil {
// 		return "Not Found"
// 	} else {return "Found"}	
// }

func ConnectDB()(*sql.DB){
	var dbConfig Config

	filename, _ := filepath.Abs("../config/mysql.yml")
	yamlFile, _ := ioutil.ReadFile(filename)
	yamlErr := yaml.Unmarshal(yamlFile, &dbConfig)
	if yamlErr != nil {
			panic(yamlErr)
	}

	connectURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)


	client, clientErr := sql.Open("mysql", connectURL)
	
	if clientErr != nil {
			log.Fatal(clientErr)
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