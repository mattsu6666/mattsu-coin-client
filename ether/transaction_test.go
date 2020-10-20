package ether

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"math/big"
	"reflect"
	"testing"
)

type MattsuCoinAPIMock struct {
	mock.Mock
}

func (mock *MattsuCoinAPIMock) Burn(to common.Address, amount *big.Int) (Transaction, error) {
	args := mock.Called(to, amount)
	return args.Get(0).(Transaction), args.Error(1)
}

func (mock *MattsuCoinAPIMock) Mint(to common.Address, amount *big.Int) (Transaction, error) {
	args := mock.Called(to, amount)
	return args.Get(0).(Transaction), args.Error(1)
}

func (mock *MattsuCoinAPIMock) BuyLatestEthUsd() (Transaction, error) {
	args := mock.Called()
	return args.Get(0).(Transaction), args.Error(1)
}

func (mock *MattsuCoinAPIMock) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	args := mock.Called(ctx, q, ch)
	return args.Get(0).(ethereum.Subscription), args.Error(1)
}

type TransactionMock struct {
	mock.Mock
}

func (mock *TransactionMock) Gas() uint64         { return mock.Called().Get(0).(uint64) }
func (mock *TransactionMock) GasPrice() *big.Int  { return mock.Called().Get(0).(*big.Int) }
func (mock *TransactionMock) Nonce() uint64       { return mock.Called().Get(0).(uint64) }
func (mock *TransactionMock) To() *common.Address { return mock.Called().Get(0).(*common.Address) }

func TestBurn(t *testing.T) {
	apiMock := new(MattsuCoinAPIMock)
	transactionMock := new(TransactionMock)
	type args struct {
		to     string
		amount int64
		api    MattsuCoinAPI
	}
	tests := []struct {
		name    string
		args    args
		want    Transaction
		wantErr bool
	}{
		{"return transaction", args{"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1, apiMock}, transactionMock, false},
		{"return error", args{"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", 1, apiMock}, nil, true},
	}
	apiMock.On("Burn", common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), big.NewInt(1)).Return(transactionMock, nil)
	apiMock.On("Burn", common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"), big.NewInt(1)).Return(transactionMock, errors.New("err!"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Burn(tt.args.to, tt.args.amount, tt.args.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("Burn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Burn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMint(t *testing.T) {
	apiMock := new(MattsuCoinAPIMock)
	transactionMock := new(TransactionMock)
	type args struct {
		to     string
		amount int64
		api    MattsuCoinAPI
	}
	tests := []struct {
		name    string
		args    args
		want    Transaction
		wantErr bool
	}{
		{"return transaction", args{"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1, apiMock}, transactionMock, false},
		{"return error", args{"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", 1, apiMock}, nil, true},
	}
	apiMock.On("Mint", common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), big.NewInt(1)).Return(transactionMock, nil)
	apiMock.On("Mint", common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"), big.NewInt(1)).Return(transactionMock, errors.New("err!"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mint(tt.args.to, tt.args.amount, tt.args.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
