package BLC
import (
    "crypto/sha256"
    "encoding/gob"
    "fmt"
    "github.com/Nik-U/pbc"
    "os"
)
var  BLS_sys   *BLS
type  BLS  struct{
    bls_pairing  *pbc.Pairing
    bls_g        *pbc.Element
}
type BLS_byte  struct{
    Params       string
    G           []byte
}
var blsfilePath="./sys/bls.gob"
func save_bls_sys_into_lib(){
    params := pbc.GenerateA(160, 512)
    para_byte:=params.String()
    g_byte:=params.NewPairing().NewG2().Rand().Bytes()
    bls_byte:=BLS_byte{para_byte,g_byte }
    blsSaveToFilebyte(bls_byte)
}
func BLS_start()*BLS{
    blsbyte_:=blsLoadFromFile()
    pairing, _ := pbc.NewPairingFromString(blsbyte_.Params)
    g := pairing.NewG2().SetBytes(blsbyte_.G)
    bls:=&BLS{pairing,g}
    return   bls
}
func  blsSaveToFilebyte(bls BLS_byte) {
    pa := &bls
    file, _ := os.OpenFile(blsfilePath, os.O_CREATE|os.O_WRONLY, 0666)
    defer file.Close()
    enc := gob.NewEncoder(file)
    enc.Encode(pa)
}
func blsLoadFromFile(  )*BLS_byte {
    file, _ := os.Open(blsfilePath)
    defer file.Close()
    var pa BLS_byte
    dec := gob.NewDecoder(file)
    dec.Decode(&pa)
    return &pa
}
func  generate_pairing()(*BLS){
    _params := pbc.GenerateA(160, 512)
    _pairing := _params.NewPairing()
    BLS_G:=_pairing.NewG2().Rand()
    bls:=&BLS{_pairing,BLS_G}
    return   bls

}
func  generate_bls_keypair(_pairing  *pbc.Pairing)(*pbc.Element,*pbc.Element){
    privKey := _pairing.NewZr().Rand()
    //g := _pairing.NewG2().Rand()
    pubKey := _pairing.NewG2().PowZn(BLS_sys.bls_g, privKey)
    return  privKey,pubKey
}
func  bls_signature(pairing  *pbc.Pairing,message  []byte,privkey *pbc.Element) *pbc.Element{
    h := pairing.NewG1().SetFromStringHash(string(message), sha256.New())
    signature := pairing.NewG2().PowZn(h, privkey)
    return  signature
}
func  bls_verify(pairing  *pbc.Pairing,g  *pbc.Element,message  []byte,pubkey  *pbc.Element,signature  *pbc.Element)bool{
    h := pairing.NewG1().SetFromStringHash(string(message), sha256.New())
    temp1 := pairing.NewGT().Pair(h, pubkey)
    temp2 := pairing.NewGT().Pair(signature, g)
    if !temp1.Equals(temp2) {
        fmt.Println("*BUG* Signature check failed *BUG*")
        return false
    } else {
        return  true
    }
}



















