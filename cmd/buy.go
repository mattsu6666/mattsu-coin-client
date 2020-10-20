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

// buyCmd represents the burn command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "buy eth/usd rate",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := ether.NewMattsuCoinClientConfig(cfgFile)
		api := ether.NewMattsuCoinAPI(cfg)
		rate, err := ether.Buy(cfg.MattsuCoinAddress, api)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rate)
	},
}

func init() {
	rootCmd.AddCommand(buyCmd)
}
