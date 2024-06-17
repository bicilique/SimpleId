package blockchains

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitConnection(networkRPC string) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connected")
	return client, err
}
