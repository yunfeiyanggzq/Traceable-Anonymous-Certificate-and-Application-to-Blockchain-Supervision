package BLC
import (
	"fmt"
	"log"
)
func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	wallets, err :=NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}



	wallet := wallets.GetWallet(from)

	cert_X:=CA_sys.pairing.NewG1().SetBytes(wallet.Random_cert_x)
	cert_Y:=CA_sys.pairing.NewG1().SetBytes(wallet.Random_cert_y)
	cert_Z:=CA_sys.pairing.NewG1().SetBytes(wallet.Random_cert_z)
	rand_cert:=RAND_CERTIFICATION{cert_X,cert_Y,cert_Z}
	varify:=MINER_verify_rand_cert(CA_sys.pairing,CA_sys.ca_pubkey,&rand_cert)
	if  varify!=true{
		log.Panic("rand_cert  verify   failed,please  register in  CA")
	}else{
		fmt.Println("rand_cert  in  tx   verify    success")
	}



	//MINER_verify_rand_cert(wallet.pairing ,wallet.CA_pubkey ,wallet.rand_cert  )
	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)
	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}
		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}
	fmt.Println("Success!")
}
