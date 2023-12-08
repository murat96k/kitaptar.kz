package config

import "time"

type Config struct {
	HttpServer HttpServer  `mapstructure:"HttpServer"`
	GrpcServer GrpcServer  `mapstructure:"GrpcServer"`
	Database   DBConfig    `mapstructure:"Database"`
	Auth       Auth        `mapstructure:"Auth"`
	Redis      RedisConfig `mapstructure:"Redis"`
	SMTP       SMTPConfig  `mapstructure:"SMTP"`
	Kafka      Kafka       `mapstructure:"Kafka"`
}

type HttpServer struct {
	Port            string        `mapstructure:"port"`
	Timeout         time.Duration `mapstructure:"timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type Auth struct {
	PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive        time.Duration `mapstructure:"TimeToLive"`
}

type GrpcServer struct {
	Port string `mapstructure:"Port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Address        string        `mapstructure:"address"`
	DB             int           `mapstructure:"db"`
	ExpirationTime time.Duration `mapstructure:"TimeToLive"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Kafka struct {
	Brokers  []string `mapstructure:"brokers"`
	Producer Producer `mapstructure:"producer"`
	Consumer Consumer `mapstructure:"consumer"`
}

type Producer struct {
	Topic string `mapstructure:"topic"`
}

type Consumer struct {
	Topics []string `mapstructure:"topics"`
}
