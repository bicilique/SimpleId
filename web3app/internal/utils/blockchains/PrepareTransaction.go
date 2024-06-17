package blockchains

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PrepareTransaction(publicKey common.Address, privateKey *ecdsa.PrivateKey, client *ethclient.Client) *bind.TransactOpts {
	// nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// auth.GasLimit = uint64(1065160) // It can change
	auth.GasLimit = uint64(8000000) // It can change
	auth.GasPrice = gasPrice
	return auth
}
