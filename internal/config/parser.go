package config

import (
	"fmt"
	"os"
	"time"

	"github.com/benu-cloud/benu-message/rabbitmq"
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

func ParseArgs() (s ServerSettings, m rabbitmq.MessageBrokerSettings) {
	// try to load env variables if they exist
	godotenv.Load()

	var serverPort PortNumber = PortNumber(8080)
	var serverTLSCert string
	var serverTLSKey string

	var rmqhost string
	var rmqport PortNumber = PortNumber(5672)
	var rmqvHost string
	var rmqusername string
	var rmqpassword string
	var rmqpublishTimeoutSeconds uint

	flag.Var(&serverPort, "httpport", "HTTP server hosting port, Should be in the range 0-65535.")
	flag.StringVar(&serverTLSCert, "httpcert", "", "TLS cert file (optional - must also provide key file).")
	flag.StringVar(&serverTLSKey, "httpkey", "", "TLS key file (optional - must also provide cert file).")

	flag.StringVar(&rmqhost, "rmqhost", "localhost", "RabbitMQ message broker host.")
	flag.Var(&rmqport, "rmqport", "RabbitMQ message broker port. Should be in the range 0-65535.")
	flag.StringVar(&rmqvHost, "rmqvhost", "", "RabbitMQ virtual host.")
	flag.StringVar(&rmqusername, "rmqusername", "", "RabbitMQ username (required).")
	flag.StringVar(&rmqpassword, "rmqpassword", "", "RabbitMQ password (required).")
	flag.UintVar(&rmqpublishTimeoutSeconds, "rmqtimeout", 5, "RabbitMQ publish timeout in seconds")

	flag.Parse()

	// check required fields
	if (serverTLSCert == "" && serverTLSKey != "") || (serverTLSKey == "" && serverTLSCert != "") {
		fmt.Println("TLS cert and key files both need to be provided.")
		flag.Usage()
		os.Exit(1)
	}
	if rmqusername == "" {
		fmt.Println("Error: the flag -rmqusername is required.")
		flag.Usage()
		os.Exit(1)
	}
	if rmqpassword == "" {
		fmt.Println("Error: the flag -rmqpassword is required.")
		flag.Usage()
		os.Exit(1)
	}

	s.Port = serverPort
	s.TLSCertFile = serverTLSCert
	s.TLSKeyFile = serverTLSKey

	m.Host = rmqhost
	m.Password = rmqpassword
	m.Port = uint(rmqport)
	m.Username = rmqusername
	m.VHost = rmqvHost
	m.PublishTimeout = time.Second * time.Duration(rmqpublishTimeoutSeconds)

	return
}
