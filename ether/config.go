package ether

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MattsuCoinClientConfig struct {
	PrivateKey        string `yaml:"privateKey"`
	KovanProjectId    string `yaml:"kovanProjectId"`
	MattsuCoinAddress string `yaml:"mattsuCoinAddress"`
	Gas               int    `yaml:"gas"`
	GasLimit          uint64 `yaml:"gasLimit"`
}

func NewMattsuCoinClientConfig(filePath string) *MattsuCoinClientConfig {
	config := MattsuCoinClientConfig{}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &config
}
