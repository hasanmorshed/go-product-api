package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configarations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func LoadConfig() {
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println("Failed to load the end file: ",err)
		os.Exit(1)
	}
	version:= os.Getenv("VERSION")
	if version==""{
		fmt.Println("Version  Is Required")
		os.Exit(1)
	}
	serviceName:= os.Getenv("SERVICE_NAME")
	if serviceName==""{
		fmt.Println("Service Name Is Required")
		os.Exit(1)
	}
	httpPort:= os.Getenv("HTTP_PORT")
	if httpPort==""{
		fmt.Println("HTTP port is required")
		os.Exit(1)
	}
	port,err:=strconv.ParseInt(httpPort,10,64)
	if err!=nil{
		fmt.Println("Port must be Number")
		os.Exit(1)
	}
	configarations = Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port),

	}

}

func GetConfig() Config{
	LoadConfig()
	return configarations
}