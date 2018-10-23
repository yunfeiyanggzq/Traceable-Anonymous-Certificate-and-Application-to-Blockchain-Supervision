package BLC

import (
	"fmt"
	"log"
)
func  (cli *CLI)traceuser(address , nodeID  string){
	fmt.Println("trace.......")
	if !ValidateAddress(address) {
		log.Panic("ERROR: Sender address is not valid")
	}
	wallets, err :=NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(address)
	R_x:=CA_sys.pairing.NewG1().SetBytes(wallet.Random_pubkey_x)
	R_y:=CA_sys.pairing.NewG1().SetBytes(wallet.Random_pubkey_y)
	R:= USER_RAND_PUBKEY{R_x,R_y}
	find_the_user_from_all(CA_sys,&R, len(wallets.Wallets))
}
