# How to work

**Please make sure that the golang environment currently exists and is >= 1.18.4**
```golang
go mod tidy
```

You can change the number of generated sui wallets in ```main.go```
```golang
package main

import (
	"SuiNetwork/sui"
)

func main() {
	dev := sui.NewSui()
    // nubmer of wallet [1：Infinity]
	dev.CreateWallet(10)
}

```
And you will see the following output and a wallet.json will be generated locally
```golang
> go run main.go

================== 正在创建Sui钱包 ==================
[+]:0x32495cea7a4dc3e05adbdc0d80417e2ceefd5f09 
[+]:0x100270dd578c09be59816325089132c380a71536 
[+]:0x1673f3376df16fdb4b71df3f8386a1ec2edc62c8 
[+]:0x5bf57e950018c0d7552b6ed4bb49df32a6d7a5ab 
[+]:0x78c4749d84ae22075e7d62c8709c8ac4e82c0116 
[+]:0x5833a110ab3c1df598ee2ec29af511b22f5540f6 
[+]:0x54a557139b45b23a9074156a3022a9e99b2b4703 
[+]:0x8adfa2940078029a5eeb126e950f831c7b7a6836 
[+]:0xe5cda58c4f28c5a642fc1ee1c86040643c1c4f0d 
======== Compile created Sui Wallet (10) ============= 
```
&nbsp;

# At the end
Please observe the rules of sui-network and do not abuse it.
