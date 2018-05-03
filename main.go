package main

import (
	"fmt"

	bip32 "github.com/FactomProject/go-bip32"
	bip44 "github.com/FactomProject/go-bip44"
	//bip32 "github.com/tyler-smith/go-bip32"
	//bip39 "github.com/tyler-smith/go-bip39"
)

func main() {
	mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"
	key, err := bip44.NewKeyFromMnemonic(mnemonic, bip44.TypeBitcoin, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("BIP44 Master key: ", key.String())
}

//func main() {

//// get the words
//reader := bufio.NewReader(os.Stdin)
//fmt.Print("Enter 23 words seperated by spaces:\n")
//text, err := reader.ReadString('\n')
//if err != nil {
//panic(err)
//}
//text = strings.TrimRight(text, "\n") // trim the enter
//words := strings.Split(text, " ")

//if len(words) != 23 {
//panic("I only take 23 words")
//}

//var bitsStr string
//for _, word := range words {
//i := bip39.ReverseWordMap[word]
//if i == 0 && word != "abandon" {
//panic(fmt.Sprintf("bad word: %v\n", word))
//}
//wordBits := strconv.FormatInt(int64(i), 2)
//wordBits = fmt.Sprintf("%011v", wordBits)
//bitsStr += wordBits
//}

//// add three bits of "entropy" lawl
////    32 * 8 - len(bitsStr) = 3
//checksumEntropys := []string{"000", "001", "010", "100", "110", "011", "101", "111"}
//for _, checksumEntropy := range checksumEntropys {
//fullBitsStr := bitsStr + checksumEntropy

//// add a space to the bits every 8 characters
//// to process the string bits to actual bytes
//var spaced string
//for i, s := range fullBitsStr {
//spaced += string(s)
//if (i+1)%8 == 0 {
//spaced += " "
//}
//}

//// process the string bits to bytes
//entropy := make([]byte, 32)
//for i, s := range strings.Fields(spaced) {
//n, _ := strconv.ParseUint(s, 2, 8)
//b := byte(n)
//entropy[i] = b
//}

//// Generate a mnemonic for memorization or user-friendly seeds
//mnemonic, _ := bip39.NewMnemonic(entropy)

//seed := bip39.NewSeed(mnemonic, "")

//masterKey, _ := bip32.NewMasterKey(seed)
//publicKey := masterKey.PublicKey()

//// Display mnemonic and keys
//fmt.Println(mnemonic)
//fmt.Println("Master private key: ", masterKey)
//fmt.Println("Master public key: ", publicKey)

//key, err := bip44.NewKeyFromMnemonic(mnemonic, 0, 0, 0, 0)
//if err != nil {
//panic(err)
//}
//fmt.Println("BIP44 Master key: ", key.String())
//}
//}
