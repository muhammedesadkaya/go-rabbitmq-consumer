package config

type ApplicationConfig struct {
	RabbitMQ   RabbitConfig        `yaml:"rabbitmq"`
	PostgreSQL PostgreSQLSqlConfig `yaml:"postgresql"`
	MongoDB    MongoDBConfig       `yaml:"rabbitmq"`
}

type RabbitConfig struct {
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Host     []string `yaml:"host"`
	Port     int      `yaml:"port"`
}

type PostgreSQLSqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

type MongoDBConfig struct {
	Address string `yaml:"address"`
}
