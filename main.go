package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"vandarGateway/payment/vandar"

	"gopkg.in/yaml.v3"
)

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type errorResponse struct {
	code    string `json:"code"`
	message string `json:"message"`
}

type config struct {
	Server server `yaml:"server"`
}

func readServerConf() (*config, error) {
	conf := new(config)
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		return nil, errors.New("error read config file")
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return nil, errors.New("error marshal config file")
	}
	return conf, nil
}

func main() {
	conf, err := readServerConf()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/get-token", vandar.GetToken)
	mux.HandleFunc("/callback", vandar.Callback)
	mux.HandleFunc("/redirect", vandar.Redirect)
	mux.HandleFunc("/transaction-detail", vandar.TransactionDetail)
	mux.HandleFunc("/verify", vandar.Verify)

	serverConfig := conf.Server.Host + ":" + conf.Server.Port
	log.Println("Start Vandar Server ", serverConfig)
	err = http.ListenAndServe(serverConfig, mux)
	log.Fatal(err)

}
