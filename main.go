package main

import (
	"fmt"
	"os"

	"github.com/AndreasBriese/bbloom"
	bip39 "github.com/tyler-smith/go-bip39"

	"github.com/rigelrozanski/common"
	"github.com/rigelrozanski/go-dicedemon/maker"
	"github.com/rigelrozanski/gowallet/wallet"
)

func main() {

	// simple word search
	// 3, 6, 9, 12, 15, 18, 21, 24 = 8 options
	// 2048 words ~ 16000 keys * 100 = 160,000 simple addresses
	// mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"

	args := os.Args[1:]
	compareCSVFile := args[0]
	compare, err := common.ReadLines(compareCSVFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("loaded file")

	//bloom
	lenBloom := float64(len(compare))
	const probCollide = 0.0001
	bf := bbloom.New(lenBloom, probCollide)
	for _, compareEl := range compare {
		bf.Add([]byte(compareEl))
	}
	fmt.Println("added to bloom")

	for mm := 0; mm < 8; mm++ {
		var mnemonicLen int8
		switch mm {
		case 0:
			mnemonicLen = int8(24)
		case 1:
			mnemonicLen = int8(21)
		case 2:
			mnemonicLen = int8(18)
		case 3:
			mnemonicLen = int8(15)
		case 4:
			mnemonicLen = int8(12)
		case 5:
			mnemonicLen = int8(9)
		case 6:
			mnemonicLen = int8(6)
		case 7:
			mnemonicLen = int8(3)
		}
		fmt.Printf("mnemonicLen: %v\n", mnemonicLen)
		for mnemonicsI := 0; mnemonicsI < 2048; mnemonicsI++ {
			mnemonics := MnemonicsForWord(mnemonicsI, mnemonicLen)
			for checksumI, mnemonic := range mnemonics {
				addrs := AddressesFromMnemonic(mnemonic, 5)
				for _, addr := range addrs {
					masterAddr := newAddrLoco(addr, mnemonicLen, int16(mnemonicsI), int16(checksumI))
					if bf.Has([]byte(masterAddr.addr)) {
						fmt.Printf("false positive: %#v\n", masterAddr)
						// now try'n find a real positive
						for _, compareEl := range compare {
							if masterAddr.addr == compareEl {
								fmt.Printf("POSITIVE: %#v\n", masterAddr)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("DONE")
}

type addrLoco struct {
	addr        string
	mnemonicLen int8
	mnemonicsI  int16
	checksumI   int16
}

func newAddrLoco(addr string, mnemonicLen int8, mnemonicsI, checksumI int16) addrLoco {
	return addrLoco{
		addr:        addr,
		mnemonicLen: mnemonicLen,
		mnemonicsI:  mnemonicsI,
		checksumI:   checksumI,
	}
}

// get bip44 addresses from a mnemonic
func AddressesFromMnemonic(mnemonic string, num uint32) []string {

	addresses := make([]string, num)
	seed := bip39.NewSeed(mnemonic, "")
	walletAcc, err := wallet.NewWalletAccountFromSeed(seed)
	if err != nil {
		panic(err)
	}
	wallets, err := walletAcc.GenerateWallets(0, num)
	if err != nil {
		panic(err)
	}
	for i, wallet := range wallets {
		addresses[i] = wallet.Address
	}
	return addresses
}

// get all the mnemonics for single word address
func MnemonicsForWord(index int, len int8) []string {

	word := bip39.EnglishWordList[index]

	words := make([]string, len-1)
	for i := int8(0); i < len-1; i++ {
		words[i] = word
	}

	return maker.PartialMnemonicToAllMnemonic(words)
}
