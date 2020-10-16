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

func Burn(to string, amount int64, cfg MattsuCoinClientConfig) (*types.Transaction, error) {
	_to := common.HexToAddress(to)
	client, instance, err := clientAndInstance(cfg.KovanProjectId, common.HexToAddress(cfg.MattsuCoinAddress))

	auth, err := auth(client, cfg)
	if err != nil {
		return nil, err
	}

	tx, err := instance.Burn(auth, _to, big.NewInt(amount))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Mint(to string, amount int64, cfg MattsuCoinClientConfig) (*types.Transaction, error) {
	_to := common.HexToAddress(to)
	client, instance, err := clientAndInstance(cfg.KovanProjectId, common.HexToAddress(cfg.MattsuCoinAddress))

	auth, err := auth(client, cfg)
	if err != nil {
		return nil, err
	}

	tx, err := instance.Mint(auth, _to, big.NewInt(amount))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Buy(cfg MattsuCoinClientConfig) (*big.Int, error) {
	client, instance, err := clientAndInstance(cfg.KovanProjectId, common.HexToAddress(cfg.MattsuCoinAddress))

	auth, err := auth(client, cfg)
	if err != nil {
		return big.NewInt(0), err
	}

	_, err = instance.BuyLatestEthUsd(auth)
	if err != nil {
		return big.NewInt(0), err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(cfg.MattsuCoinAddress)},
	}

	logs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, logs)
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

func clientAndInstance(projectId string, address common.Address) (*ethclient.Client, *MattsuCoin, error) {
	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/" + projectId)
	if err != nil {
		return nil, nil, err
	}

	instance, err := NewMattsuCoin(address, client)
	if err != nil {
		return nil, nil, err
	}

	return client, instance, nil
}

func auth(client *ethclient.Client, cfg MattsuCoinClientConfig) (*bind.TransactOpts, error) {
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
