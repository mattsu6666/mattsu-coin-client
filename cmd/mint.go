/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/mattsu6666/mattsu-coin-client/ether"
	"github.com/spf13/cobra"
	"log"
)

var mintTo string
var mintAmount int64

// mintCmd represents the mint command
var mintCmd = &cobra.Command{
	Use:   "mint",
	Short: "min MattsuCoin",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := ether.NewMattsuCoinClientConfig(cfgFile)
		api := ether.NewMattsuCoinAPI(cfg)
		tx, err := ether.Mint(mintTo, mintAmount, api)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("nonce: %d, gas: %d, gasPrice: %d, to: %s\n", tx.Nonce(), tx.Gas(), tx.GasPrice(), tx.To().Hex())
	},
}

func init() {
	rootCmd.AddCommand(mintCmd)
	mintCmd.Flags().StringVar(&mintTo, "to", "", "to address. ex. 0xabcd...")
	_ = mintCmd.MarkFlagRequired("to")
	mintCmd.Flags().Int64Var(&mintAmount, "amount", 0, "amount. ex. 1000")
	_ = mintCmd.MarkFlagRequired("amount")
}
