package main

import (
	"SuiNetwork/sui"
)

func main() {
	dev := sui.NewSui()
	dev.CreateWallet(10)
}
