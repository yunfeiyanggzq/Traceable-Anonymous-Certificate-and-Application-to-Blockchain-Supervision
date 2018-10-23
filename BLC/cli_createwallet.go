package BLC

import (
	"fmt"
)
func (cli *CLI) createWallet(nodeID string) {

	wallets, _ := NewWallets(nodeID)

	address := wallets.CreateWallet(len(wallets.Wallets))
	wallets.SaveToFile(nodeID)
	fmt.Printf("Your new address: %s\n", address)

}
