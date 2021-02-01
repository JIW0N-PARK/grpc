package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"

	pb "../proto/test"
	"google.golang.org/grpc"

	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gopkg.in/yaml.v2"
)

type server struct {
	pb.UnimplementedWebServer
	Database *Database
}

type User struct {
	Id   int32 `gorm:"primaryKey"`
	Name string
	Age  int32
}

type Database struct {
	DB *gorm.DB
}

type Config struct {
	DBName   string `yaml:"dbName"`
	DBType   string `yaml:"dbType"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Type     string `yaml:"type"`
}

func (c *Config) setConfig(file string) {
	filename, _ := filepath.Abs(file)
	yamlFile, _ := ioutil.ReadFile(filename)
	yamlErr := yaml.Unmarshal(yamlFile, &c)
	if yamlErr != nil {
		panic(yamlErr)
	}
}

const (
	port = ":50051"
)

func (s server) Register(ctx context.Context, in *pb.RegisterReq) (*pb.Response, error) {
	log.Printf("--Register--")
	log.Printf("Name=%s\n", in.GetName())
	log.Printf("Age=%d\n", in.GetAge())

	res := InsertPost(s.Database.DB, in.GetName(), in.GetAge())
	return &pb.Response{Message: res}, nil
}

func (s server) Check(ctx context.Context, in *pb.CheckReq) (*pb.Response, error) {
	res := GetPost(s.Database.DB, in.GetName(), &User{})
	log.Printf("--Retrive--")
	log.Printf("Name=%s\n", in.GetName())
	log.Printf("Result - %s\n", res)
	return &pb.Response{Message: res}, nil
}

func InsertPost(db *gorm.DB, name string, age int32) string {
	var count int32
	db.Model(&User{}).Count(&count)
	user := User{Id: count + 1, Name: name, Age: age}

	result := db.Create(&user)
	if result.RowsAffected == 1 {
		return "Complete Insert!"
	} else {
		return "Failed.."
	}
	//db.Exec("INSERT INTO grpc VALUES (?, ?, ?)", 5, name, age)
}

func GetPost(db *gorm.DB, name string, user *User) string {
	result := db.Where("name = ?", name).Find(&user)
	// db.Query("SELECT * FROM GRPC WHERE NAME = ?", name)

	if result.RowsAffected == 1 {
		return name + " is Found"
	} else {
		return name + " is Not Found"
	}
}

func dbConnect(config *Config) *gorm.DB {

	connectURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(config.DBType, connectURL)

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var config Config
	config.setConfig("../config/mysql.yml")

	db := new(Database)
	db.DB = dbConnect(&config)
	defer db.DB.Close()

	s := grpc.NewServer()
	pb.RegisterWebServer(s, &server{Database: db})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
