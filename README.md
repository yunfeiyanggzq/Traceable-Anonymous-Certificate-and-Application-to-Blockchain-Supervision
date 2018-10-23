# Traceable-Anonymous-Certificate-and-Application-to-Blockchain-Supervision
a   Traceable  Blockchain  code   in   gocode  using   TAC

# How to run 
## First:Install the  PBC  lib

 This package must be compiled using cgo. It also requires the installation of GMP and PBC. During the build process, this package will attempt to include <gmp.h> and <pbc/pbc.h>, and then dynamically link to GMP and PBC.

Most systems include a package for GMP. To install GMP in Debian / Ubuntu:

`sudo apt-get install libgmp-dev`

For an RPM installation with YUM:

`sudo yum install gmp-devel`

For installation with Fink (http://www.finkproject.org/) on Mac OS X:

`sudo fink install gmp gmp-shlibs`

For more information or to compile from source, visit https://gmplib.org/

To install the PBC library, download the appropriate files for your system from https://crypto.stanford.edu/pbc/download.html. PBC has three dependencies: the gcc compiler, flex (http://flex.sourceforge.net/), and bison (https://www.gnu.org/software/bison/). See the respective sites for installation instructions. Most distributions include packages for these libraries. For example, in Debian / Ubuntu:

`sudo apt-get install build-essential flex bison`

The PBC source can be compiled and installed using the usual GNU Build System:
```
./configure
make
sudo make install
```

After installing, you may need to rebuild the search path for libraries:

`sudo ldconfig`

It is possible to install the package on Windows through the use of MinGW and MSYS. MSYS is required for installing PBC, while GMP can be installed through a package. Based on your MinGW installation, you may need to add `"-I/usr/local/include"` to `CPPFLAGS `and `"-L/usr/local/lib"` to LDFLAGS when building PBC. Likewise, you may need to add these options to` CGO_CPPFLAGS `and` CGO_LDFLAGS` when installing this package. 

then  install the PBC go  lib

`go  get  github.com/Nik-U/pbc`

## Second:Install the  bolt  lib
put  in   terminal  
`go  get   github.com/boltdb/bolt`
# Useage
```
yu@ubuntu:~$ cd  go
yu@ubuntu:~/go$ cd src
yu@ubuntu:~/go/src$ go  build  coin/main.go
yu@ubuntu:~/go/src$ ./main
==================================================================================================
****************************************Usage:****************************************************
==================================================================================================
  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS
  createwallet  -name  NAME  - Generates a new key-pair and saves it into the wallet file
  getbalance -address ADDRESS - Get balance of ADDRESS
  listaddresses - Lists all addresses from the wallet file
  printchain - Print all the blocks of the blockchain
  reindexutxo - Rebuilds the UTXO set
  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.
  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining
  trace  -userAddress  ADDRESS -trace  the  user  register  in CA
==================================================================================================

```
# Run
```

yu@ubuntu:~/go/src$ ./main createwallet
NODE_ID env. var is not set!yu@ubuntu:~/go/src$ 
yu@ubuntu:~/go/src$ export  NODE_ID=3000
yu@ubuntu:~/go/src$ ./main createwallet  -name  tom
your  name  in CA:tom
Your new address: 1DPEPgMmURDgWNVkzs739tYB2QheapujT8
yu@ubuntu:~/go/src$ ./main createwallet  -name  Tony
your  name  in CA:Tony
Your new address: 1GXgEgLX8VctNTZAZsv5PpH1tAi7iUW2e4
yu@ubuntu:~/go/src$ ./main createblockchain -address 1GXgEgLX8VctNTZAZsv5PpH1tAi7iUW2e4
0000a84414844e388658849b116f93902a556892d36e157b00949f6b0080891d
Done!
yu@ubuntu:~/go/src$ ./main send -from 1GXgEgLX8VctNTZAZsv5PpH1tAi7iUW2e4  -to 1DPEPgMmURDgWNVkzs739tYB2QheapujT8 -amount 4 -mine
rand_cert  in  tx   verify    success
signature  in  tx   verify    success
0000d521d7bac20c264e9029db39f92adb77c29d5329cfb08cdb1f20d559a1fb
Success!
yu@ubuntu:~/go/src$ ./main trace  -userAddress   1GXgEgLX8VctNTZAZsv5PpH1tAi7iUW2e4
trace....................................................
it  is   the  2th  user  registered   in   CA ,tom. 
```
