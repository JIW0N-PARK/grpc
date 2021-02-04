type dbType struct {
	mysql  *dbConfig
	oracle *dbConfig
}

type dbConfig struct {
	DBName       string `yaml:"dbName"`
	DBType       string `yaml:"dbType"`
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	MaxIdleConns int32  `yaml:"maxIdleConns"`
	MaxOpenConns int32  `yaml:"maxOpenConns"`
}