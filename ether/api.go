package ether

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type MattsuCoinAPI interface {
	Burn(to common.Address, amount *big.Int) (Transaction, error)
	Mint(to common.Address, amount *big.Int) (Transaction, error)
	BuyLatestEthUsd() (Transaction, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

type MattsuCoinAPIImpl struct {
	instance *MattsuCoin
	client   *ethclient.Client
	auth     *bind.TransactOpts
}

func (api *MattsuCoinAPIImpl) Burn(to common.Address, amount *big.Int) (Transaction, error) {
	tx, err := api.instance.Burn(api.auth, to, amount)
	return &TransactionImpl{tx}, err
}

func (api *MattsuCoinAPIImpl) Mint(to common.Address, amount *big.Int) (Transaction, error) {
	tx, err := api.instance.Mint(api.auth, to, amount)
	return &TransactionImpl{tx}, err

}

func (api *MattsuCoinAPIImpl) BuyLatestEthUsd() (Transaction, error) {
	tx, err := api.instance.BuyLatestEthUsd(api.auth)
	return &TransactionImpl{tx}, err
}

func (api *MattsuCoinAPIImpl) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return api.client.SubscribeFilterLogs(ctx, q, ch)
}

func NewMattsuCoinAPI(cfg *MattsuCoinClientConfig) *MattsuCoinAPIImpl {
	client, instance, err := ClientAndInstance(cfg)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := Auth(client, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &MattsuCoinAPIImpl{
		instance: instance,
		client:   client,
		auth:     auth,
	}
}
