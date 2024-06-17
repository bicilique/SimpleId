package blockchains

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type Configurations struct {
	Nonce    uint64
	GasPrice *big.Int
	ChainID  *big.Int
	GasLimit uint64
}

func TransferOptions(publicKey common.Address, client *ethclient.Client) Configurations {
	var configure Configurations

	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	configure.Nonce = nonce
	configure.GasPrice = gasPrice
	configure.ChainID = chainID
	configure.GasLimit = uint64(21000)
	return configure
}

func ToWei(iAmount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iAmount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func Transfer(fromPublicAddress common.Address, fromPrivateAddress *ecdsa.PrivateKey, toPublicAddress common.Address, amount string, client *ethclient.Client) {
	weiAmount := ToWei(amount, 18)

	transactionConfigs := TransferOptions(fromPublicAddress, client)

	var data []byte
	tx := types.NewTransaction(transactionConfigs.Nonce, toPublicAddress, weiAmount, transactionConfigs.GasLimit, transactionConfigs.GasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(transactionConfigs.ChainID), fromPrivateAddress)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
