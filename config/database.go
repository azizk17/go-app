package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	// Driver            string        `mapstructure:"DRIVER"`
	// Source            string        `mapstructure:"SOURCE"`
	// ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	// TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	// AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

	Driver      string `mapstructure:"db_driver" yaml:"db_driver" env:"DB_DRIVER"`
	Host        string `mapstructure:"db_host" yaml:"db_host" env:"DB_HOST"`
	Username    string `mapstructure:"db_username" yaml:"db_username" env:"DB_USER"`
	Password    string `mapstructure:"db_password" yaml:"db_password" env:"DB_PASS"`
	DBName      string `mapstructure:"db_name" yaml:"db_name" env:"DB_NAME"`
	Port        int    `mapstructure:"db_port" yaml:"db_port" env:"DB_PORT"`
	Connections int    `mapstructure:"db_connections" yaml:"db_connections" env:"DB_CONNECTIONS"`
}

func (d *DB) Setup() error {

	color.New(color.FgHiMagenta).Printf("DB HOST: %v\n", d.Host)
	var connectionString string
	switch d.Driver {
	case "postgres":
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", d.Host, d.Port, d.Username, d.DBName, d.Password)
	default:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.DBName)
	}
	fmt.Printf("DB DRIVER: %v \n", d.Driver)
	_, err := sql.Open(d.Driver, connectionString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// store := db.NewStore(conn)

	return nil
}
