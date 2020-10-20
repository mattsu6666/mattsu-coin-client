package ether

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type Transaction interface {
	Gas() uint64
	GasPrice() *big.Int
	Nonce() uint64
	To() *common.Address
}

type TransactionImpl struct {
	tx *types.Transaction
}

func (tx *TransactionImpl) Gas() uint64         { return tx.tx.Gas() }
func (tx *TransactionImpl) GasPrice() *big.Int  { return tx.tx.GasPrice() }
func (tx *TransactionImpl) Nonce() uint64       { return tx.tx.Nonce() }
func (tx *TransactionImpl) To() *common.Address { return tx.tx.To() }

func Burn(to string, amount int64, api MattsuCoinAPI) (Transaction, error) {
	_to := common.HexToAddress(to)
	tx, err := api.Burn(_to, big.NewInt(amount))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Mint(to string, amount int64, api MattsuCoinAPI) (Transaction, error) {
	_to := common.HexToAddress(to)
	tx, err := api.Mint(_to, big.NewInt(amount))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Buy(address string, api MattsuCoinAPI) (*big.Int, error) {
	_, err := api.BuyLatestEthUsd()
	if err != nil {
		return big.NewInt(0), err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(address)},
	}

	logs := make(chan types.Log)
	_, err = api.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	event := struct {
		Buyer  common.Address
		EthUsd *big.Int
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(MattsuCoinABI))
	if err != nil {
		log.Fatal(err)
	}

	<-logs // 不要なイベントを1つ読み飛ばす
	err = contractAbi.Unpack(&event, "BuyEthUsd", (<-logs).Data)
	if err != nil {
		log.Fatal(err)
	}

	return event.EthUsd, nil
}

func ClientAndInstance(cfg *MattsuCoinClientConfig) (*ethclient.Client, *MattsuCoin, error) {
	_address := common.HexToAddress(cfg.MattsuCoinAddress)
	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/" + cfg.KovanProjectId)
	if err != nil {
		return nil, nil, err
	}

	instance, err := NewMattsuCoin(_address, client)
	if err != nil {
		return nil, nil, err
	}

	return client, instance, nil
}

func Auth(client *ethclient.Client, cfg *MattsuCoinClientConfig) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = cfg.GasLimit
	auth.GasPrice = gasPrice
	return auth, nil
}
