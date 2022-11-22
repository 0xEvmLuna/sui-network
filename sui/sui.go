package sui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/coming-chat/go-sui/account"
	"github.com/coming-chat/go-sui/client"
	"github.com/tyler-smith/go-bip39"
)

type SuiNetwork struct {
	Rpc     *client.Client
	Wallets []*wallet
}

type wallet struct {
	Address  string `json:"address"`
	Mnemonic string `json:"mnemonic"`
	Private  string `json:"private"`
	Public   string `json:"public"`
}

func NewSui() *SuiNetwork {
	cli, err := client.Dial("https://fullnode.devnet.sui.io:443")
	if err != nil {
		log.Fatal(err)
	}

	return &SuiNetwork{
		Rpc: cli,
	}
}

func (s *SuiNetwork) CreateWallet(amount int) {
	ch := make(chan *wallet, amount)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(act chan<- *wallet) {
		defer wg.Done()

		var i int = 1
		fmt.Println("================== 正在创建Sui钱包 ==================")

		for i < amount {
			entropy, _ := bip39.NewEntropy(128)
			mnemonic, _ := bip39.NewMnemonic(entropy)
			account, _ := account.NewAccountWithMnemonic(mnemonic)
			fmt.Printf("[+]:%s \n", account.Address)

			act <- &wallet{
				Mnemonic: mnemonic,
				Address:  account.Address,
				Private:  fmt.Sprintf("%x", account.PrivateKey[:32]),
				Public:   fmt.Sprintf("%x", account.PublicKey[:32]),
			}
			i++
		}
		close(ch)
		fmt.Printf("======== Compile created Sui Wallet (%d) ============= \n", i)
	}(ch)
	wg.Wait()

	for act := range ch {
		fmt.Printf("[address]:%s [words]: %s", act.Address, act.Mnemonic)
		s.Wallets = append(s.Wallets, act)
	}

	s.write()
}

func (s *SuiNetwork) write() {
	walletBytes, _ := json.Marshal(s.Wallets)
	for _, v := range s.Wallets {
		fmt.Printf("[Address]:%s [Mnemonic]: %s \n", v.Address, v.Mnemonic)
	}
	ioutil.WriteFile(
		"SuiWallet.json",
		walletBytes,
		0777,
	)
}
