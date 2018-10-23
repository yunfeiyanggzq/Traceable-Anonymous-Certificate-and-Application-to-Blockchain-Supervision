package BLC

import (
	"bytes"
	"crypto/sha256"
	"github.com/Nik-U/pbc"
	"golang.org/x/crypto/ripemd160"
	"log"
)
const version = byte(0x00)
const addressChecksumLen = 4
// Wallet stores private and public keys
type Wallet struct {
	PrivateKey       []byte
	PublicKey        []byte
	Random_pubkey_x  []byte
	Random_pubkey_y  []byte
	Random_cert_x    []byte
	Random_cert_y    []byte
	Random_cert_z    []byte
}

// NewWallet creates and returns a Wallet
func NewWallet(numx int) *Wallet {
	user_number:=numx
	private, public := newKeyPair()
	USER_pubkey,_,T,user_rand_pubkey,rand_cert,certification:=USER_start(CA_sys.pairing,CA_sys.ca_pubkey ,CA_sys.ca_prikey)
	CA_append_user_into_lib(T ,certification ,USER_pubkey ,user_number )
	x:=user_rand_pubkey.x.Bytes()
	y:=user_rand_pubkey.y.Bytes()
	z:=rand_cert.x.Bytes()
	u:=rand_cert.y.Bytes()
	v:=rand_cert.z.Bytes()
	wallet := Wallet{private.Bytes(), public,x,y,z,u,v}
	user_number++
	UserNumSaveToFilebyte(user_number)
	return &wallet
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))
	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (*pbc.Element, []byte) {
	privKey,pubKey:=generate_bls_keypair(BLS_sys.bls_pairing)
	pubKey_byte:=pubKey.Bytes()
	return  privKey,pubKey_byte
}
