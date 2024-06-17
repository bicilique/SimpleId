package blockchains

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WalletBalanceByBlock(address string, block int64, client *ethclient.Client) (*big.Int, *big.Float) {
	account := common.HexToAddress(address)
	blockNumber := big.NewInt(block)
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return balance, ethValue
}
