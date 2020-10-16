# mattsu-coin-client
- https://github.com/mattsu6666/mattsu-coin を利用するためのクライアントアプリ

## チュートリアル

- セットアップ

```
$ mv configs/.mattsu-coin-client.yaml.sample configs/.mattsu-coin-client.yaml
$ vim configs/.mattsu-coin-client.yaml // 適宜編集
$ make all
```

- mint

```
./mattsu-coin-client mint --to "0x950650b9eA3270494529c64695Ae91689927AC97" --amount 100 --config ./configs/.mattsu-coin-client.yaml
Using config file: ./configs/.mattsu-coin-client.yaml
nonce: 60, gas: 5000000, gasPrice: 20000000000, to: 0x950650b9eA3270494529c64695Ae91689927AC97
```

mint実行前後のbalance
```
truffle(kovan)> a.balanceOf(accounts[0]).then(v => v.toString())
'11000'
...
truffle(kovan)> a.balanceOf(accounts[0]).then(v => v.toString())
'11100'
```

- burn

```
$ ./mattsu-coin-client burn --to "0xF2b3Df99E3A87984CdE3C7fB2fF0D99Ffb9905fD" --amount 100 --config ./configs/.mattsu-coin-client.yaml
Using config file: ./configs/.mattsu-coin-client.yaml
nonce: 63, gas: 5000000, gasPrice: 20000000000, to: 0x950650b9eA3270494529c64695Ae91689927AC97
```

burn実行前後のbalance
```
truffle(kovan)> a.balanceOf(accounts[0]).then(v => v.toString())
'11100'
...
truffle(kovan)> a.balanceOf(accounts[0]).then(v => v.toString())
'11000'
```

- buy

```
$ go run main.go buy --config ./configs/.mattsu-coin-client.yaml                                                   
Using config file: ./configs/.mattsu-coin-client.yaml
37885000000
```