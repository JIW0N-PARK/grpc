package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/yaml.v2"
)

type Database struct {
	config *dbConfig
}

type dbConfig struct {
	DBName       string `yaml:"dbName"`
	DBType       string `yaml:"dbType"`
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

func dbConnect(dbtype string) *gorm.DB {
	database := initDB(dbtype)

	connectURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		database.config.User, database.config.Password, database.config.Host, database.config.Port, database.config.DBName)

	db, err := gorm.Open(dbtype, connectURL)

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(database.config.MaxIdleConns)
	db.DB().SetMaxOpenConns(database.config.MaxOpenConns)

	return db
}

func initDB(dbtype string) *Database {
	database := new(Database)
	var config dbConfig
	file := "../config/" + dbtype + ".yml"
	database.config = config.setConfig(file)
	return database
}

func (c *dbConfig) setConfig(file string) *dbConfig {
	filename, _ := filepath.Abs(file)
	yamlFile, _ := ioutil.ReadFile(filename)
	yamlErr := yaml.Unmarshal(yamlFile, &c)
	if yamlErr != nil {
		panic(yamlErr)
	}
	return c
}
