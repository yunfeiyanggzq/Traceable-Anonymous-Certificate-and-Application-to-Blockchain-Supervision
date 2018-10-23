package BLC
import (
	"log"
	"os"
	"encoding/gob"
	"fmt"
	"github.com/Nik-U/pbc"
)
type  SIGNATURE  struct{
	x    *pbc.Element
	y    *pbc.Element
}
type  CAPUBKEY    struct{
	x   *pbc.Element
	y   *pbc.Element
	z   *pbc.Element
}
type  USERPUBKEY struct{
	x   *pbc.Element
	y   *pbc.Element
}
type  CAPRIKEY  struct{
	x    *pbc.Element
	y    *pbc.Element
}
type  USER_INFO  struct{
	cert   *SIGNATURE
	userpubkey  *USERPUBKEY
	T    *pbc.Element
}
type   USER_RAND_PUBKEY struct {
	x    *pbc.Element
	y    *pbc.Element
}
type   RAND_CERTIFICATION  struct{
	x    *pbc.Element
	y    *pbc.Element
	z    *pbc.Element
}
type  CA   struct{
	pairing  *pbc.Pairing
	ca_pubkey *CAPUBKEY
	ca_prikey *CAPRIKEY
}
type  CA_byte   struct {
	Parm    string
	Pub_x   []byte
	Pub_y   []byte
	Pub_z   []byte
	Pri_x   []byte
	Pri_y   []byte
}
type USER_INFO_BYTES struct {
	Cert_x          []byte
	Cert_y          []byte
	Userpubkey_x    []byte
	Userpubkey_y    []byte
	T               []byte
}
var CA_sys    *CA
var filePath="./userinfo/user_%d.gob"
var user_number  int
var cafilePath="./sys/ca.gob"
var numfile="./sys/user_num.gob"



func CA_save(){
	params := pbc.GenerateF(160)
	parm_string:=params.String()
	pairing := params.NewPairing()
	g:= pairing.NewG2().Rand()
	sharedG := g.Bytes()
	//USER_LIB:=make(map[*pbc.Element]*USER_INFO)
	CA_pubkey,CA_prikey:=CA_produce_keypair(pairing,sharedG)
	pub_x:=CA_pubkey.x.Bytes()
	pub_y:=CA_pubkey.y.Bytes()
	pub_z:=CA_pubkey.z.Bytes()
	pri_x:=CA_prikey.x.Bytes()
	pri_y:=CA_prikey.y.Bytes()
	ca_sys_byte:=CA_byte{parm_string,pub_x,pub_y,pub_z,pri_x,pri_y}
	CASaveToFilebyte(ca_sys_byte )
}
func CA_start()(*CA){
	casysbyte:=CALoadFromFile()
	pairing, _ := pbc.NewPairingFromString(casysbyte.Parm)
	pubx:=pairing.NewG2().SetBytes(casysbyte.Pub_x)
	puby:=pairing.NewG2().SetBytes(casysbyte.Pub_y)
	pubz:=pairing.NewG2().SetBytes(casysbyte.Pub_z)
	prix:=pairing.NewZr().SetBytes(casysbyte.Pri_x)
	priy:=pairing.NewZr().SetBytes(casysbyte.Pri_y)
	pub:=&CAPUBKEY{pubx,puby,pubz}
	pri:=&CAPRIKEY{prix,priy}

	ca:=CA{pairing,pub,pri}
	return   &ca
}
func  CASaveToFilebyte(ca_sys_byte  CA_byte) {
	pa := &ca_sys_byte
	file, _ := os.OpenFile(cafilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}
func CALoadFromFile(  )*CA_byte{
	file, _ := os.Open(cafilePath)
	defer file.Close()
	var pa CA_byte
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}
func  UserNumSaveToFilebyte(x  int) {
	pa := x
	file, _ := os.OpenFile(numfile, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}
func UserNumLoadFromFile(  )int {
	file, err := os.Open(numfile)
	if  err!=nil{
		return  0
	}
	defer file.Close()
	var pa int
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	fmt.Println("***user_number_load***")
	fmt.Println(pa)
	return pa
}
func CA_append_user_into_lib(T  *pbc.Element,certification *SIGNATURE,userpubkey  *USERPUBKEY,user_num int){
	user_info:=&USER_INFO{certification,userpubkey,T}
	filePath = fmt.Sprintf(filePath,user_num)
	user_info_byte:=SET_userlib_into_byte (user_info)
	SaveToFilebyte( user_info_byte)
	filePath="./userinfo/user_%d.gob"
}
func CA_load_user_from_lib(ca_sys   *CA,user_num int)USER_INFO{
	filePath = fmt.Sprintf(filePath,user_num) //fmt.Println(filePath )
	userlib_byte:=LoadFromFile()
	userlib:=SET_userlib_into_pbc(ca_sys,userlib_byte)
	filePath="./userinfo/user_%d.gob"
	return  userlib
}
func SET_userlib_into_pbc(CA_sys  *CA,userinfo *USER_INFO_BYTES) USER_INFO{
	csertx:= CA_sys.pairing.NewG2().SetBytes(userinfo.Cert_x)
	cserty:= CA_sys.pairing.NewG2().SetBytes(userinfo.Cert_y)
	userpubkeyx:= CA_sys.pairing.NewG2().SetBytes(userinfo.Userpubkey_x)
	userpubkeyy:= CA_sys.pairing.NewG2().SetBytes(userinfo.Userpubkey_y)
	T:=CA_sys.pairing.NewG2().SetBytes(userinfo.T)
	sig:=&SIGNATURE{csertx,cserty}
	pubkey:=&USERPUBKEY{userpubkeyx,userpubkeyy}
	userinfob:=USER_INFO{sig,pubkey,T}
	return  userinfob

}
// SaveToFile saves wallets to a file
func  SaveToFilebyte(userifo_byte  USER_INFO_BYTES) {
	pa := &userifo_byte
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}
func LoadFromFile(  )*USER_INFO_BYTES {

	file, _ := os.Open(filePath)
	defer file.Close()
	var pa USER_INFO_BYTES
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}
func SET_userlib_into_byte(userinfo *USER_INFO) USER_INFO_BYTES{
	cert_x:=userinfo.cert.x.Bytes()
	cert_y:=userinfo.cert.y.Bytes()
	userpubkey_x:=userinfo.userpubkey.x.Bytes()
	userpubkey_y:=userinfo.userpubkey.y.Bytes()
	T:=userinfo.T.Bytes()
	userifo_byte:=USER_INFO_BYTES{cert_x,cert_y,userpubkey_x,userpubkey_y,T}
	return  userifo_byte
}
func  find_the_user_from_all(CA_sys *CA,user_rand_pubkey_trace  * USER_RAND_PUBKEY,x  int)bool{
	User_number:=x
	for  i:=0;i<=User_number;i++{
		userinlib:=CA_load_user_from_lib(CA_sys,i)
		isfind:=CA_trace(CA_sys.pairing,user_rand_pubkey_trace,userinlib.T,CA_sys.ca_pubkey)
		if  isfind==true{
			fmt.Printf("it  is   the  %dth  user  registered  in   CA     we   trace \n ",i+1)
			return  true
		}
		filePath="./userinfo/user_%d.gob"
	}
	fmt.Println("no  user  found  in   CA\n")
	return  false
}

func USER_start(pairing *pbc.Pairing,CA_pubkey *CAPUBKEY,CA_prikey *CAPRIKEY)(*USERPUBKEY,*pbc.Element,*pbc.Element,*USER_RAND_PUBKEY,*RAND_CERTIFICATION,*SIGNATURE){
	user_number+=1
	USER_pubkey,userprikey,T:=USER_produce_keypair(pairing,CA_pubkey)
	certification           :=CA_produce_cer(pairing , USER_pubkey,CA_prikey)
	err:=USER_verify_cert(pairing ,CA_pubkey,certification,userprikey)
	if err==false{
		log.Panic(err)
	}
	user_rand_pubkey,rand_cert:=USER_random_cert_and_pubkey(pairing ,userprikey,USER_pubkey,certification)
	return  USER_pubkey,userprikey,T,user_rand_pubkey,rand_cert,certification
}
func CA_produce_keypair(pairing *pbc.Pairing, sharedG []byte)(*CAPUBKEY,*CAPRIKEY){
	g := pairing.NewG2().SetBytes(sharedG)
	privKey_x := pairing.NewZr().Rand()
	privKey_y := pairing.NewZr().Rand()
	pubKey_y:= pairing.NewG2().Rand().PowZn(g, privKey_x)
	pubKey_z:= pairing.NewG2().Rand().PowZn(g, privKey_y)
	pubKey_x:=g
	//fmt.Println("here")
	var  mypublickey=CAPUBKEY{pubKey_x,pubKey_y,pubKey_z}
	var  myprikey=CAPRIKEY{privKey_x,privKey_y}
	//fmt.Println("CA_produce_keypair")
	return &mypublickey,&myprikey
}
func USER_produce_keypair(pairing *pbc.Pairing,capubkey  *CAPUBKEY)(*USERPUBKEY,*pbc.Element,*pbc.Element) {
	//pairing, _ := pbc.NewPairingFromString(sharedParams)
	g:=pairing.NewG1().Rand()
	aphal:=pairing.NewZr().Rand()
	Y:=pairing.NewG1().PowZn(g,aphal)
	var  userpubkey=USERPUBKEY{g,Y}
	g_:=capubkey.x
	T:=pairing.NewG2().PowZn(g_,aphal)
	//fmt.Println("USER_produce_keypair")
	return &userpubkey,aphal,T

}
func CA_produce_cer(pairing *pbc.Pairing,userpubkey *USERPUBKEY,caprikeykey *CAPRIKEY)(*SIGNATURE){
	r:=pairing.NewZr().Rand()
	data1:=pairing.NewG1().PowZn(userpubkey.x,r)
	data2_1:=pairing.NewG1().PowZn(userpubkey.y,r)
	data2_2:=pairing.NewG1().PowZn(data2_1,caprikeykey.y)
	data2_3:=pairing.NewG1().PowZn(data1,caprikeykey.x)
	data2:=pairing.NewG1().Mul(data2_2,data2_3)
	signature:=SIGNATURE{data1,data2}
	//fmt.Println("CA_produce_cer")
	return  &signature
}
func  USER_verify_cert(pairing *pbc.Pairing,capubkey *CAPUBKEY,certificate *SIGNATURE,userprikey *pbc.Element)bool{
	//fmt.Print("USER_verify_cert:")
	u:=pairing.NewG2().PowZn(capubkey.z,userprikey)
	v:=pairing.NewG2().Mul(capubkey.y,u)
	temp1:=pairing.NewGT().Pair(certificate.x,v)
	temp2:=pairing.NewGT().Pair(certificate.y,capubkey.x)
	if !temp1.Equals(temp2){
		//fmt.Println("Certification failed ")
		return  false
	} else {
		//fmt.Println("Certification right")
		return  true
	}
}
func USER_random_cert_and_pubkey(pairing *pbc.Pairing,userprikey  *pbc.Element,userpukkey  *USERPUBKEY,certification *SIGNATURE)(*USER_RAND_PUBKEY,*RAND_CERTIFICATION){
	u:=pairing.NewZr().Rand()
	g_1:=pairing.NewG1().PowZn(userpukkey.x,u)
	X_1:=pairing.NewG1().PowZn(userpukkey.y,u)
	user_rand_pubkey:=USER_RAND_PUBKEY{g_1,X_1}
	v:=pairing.NewZr().Rand()
	data_1:=pairing.NewG1().PowZn(certification.x,v)
	data_2:=pairing.NewG1().PowZn(certification.y,v)
	data_3:=pairing.NewG1().PowZn(data_1,userprikey)
	user_rand_cert:=RAND_CERTIFICATION{data_1,data_2,data_3}
	return &user_rand_pubkey,&user_rand_cert
}
func CA_trace(pairing *pbc.Pairing,user_rand_pubkey *USER_RAND_PUBKEY,T *pbc.Element,capubkey *CAPUBKEY)bool{
	temp1:=pairing.NewGT().Pair(user_rand_pubkey.y,capubkey.x)
	temp2:=pairing.NewGT().Pair(user_rand_pubkey.x,T)
	if !temp1.Equals(temp2){
		return  false
	} else {
		return  true
	}
}
func  MINER_verify_rand_cert(pairing *pbc.Pairing,ca_pubkey   *CAPUBKEY,ca_rand_cert *RAND_CERTIFICATION)bool{
	//fmt.Print("MINER_verify_rand_cert:")
	temp1:=pairing.NewGT().Pair(ca_rand_cert.x,ca_pubkey.y)
	temp2:=pairing.NewGT().Pair(ca_rand_cert.z,ca_pubkey.z)
	temp3:=pairing.NewGT().Pair(ca_rand_cert.y,ca_pubkey.x)
	temp4:=pairing.NewGT().Mul(temp1,temp2)
	if !temp3.Equals(temp4){
		log.Panic("rand_cert  verify  failed")
		return  false
	} else {
		//fmt.Println("rand_cert  verify  success")
		return  true
	}
}

