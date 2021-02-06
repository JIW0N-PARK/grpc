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

func DBInit(dbtype string) *gorm.DB {
	var config dbConfig
	config.setConfig(dbtype)

	db := connect(&config)
	return db
}

func connect(config *dbConfig) *gorm.DB {
	connectURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(config.DBType, connectURL)

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	return db
}

func (c *dbConfig) setConfig(dbtype string) {
	file := "../usersys/config/"+dbtype+".yml"
	filename, _ := filepath.Abs(file)
	yamlFile, _ := ioutil.ReadFile(filename)
	yamlErr := yaml.Unmarshal(yamlFile, &c)
	if yamlErr != nil {
		panic(yamlErr)
	}
}
