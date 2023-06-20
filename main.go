package main

import (
	"fmt"

	evm "github.com/TafveezA/EVM/EVM"
)

func main() {
	contract()
}

func contract() []byte {
	data := []byte{0x03, 0x0a, 0x02, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x046, 0x0c, 0x0d, 0x03, 0x0a, 0x0f}
	pushFoo := []byte{}
	data = append(data, pushFoo...)
	contractState := evm.NewState()
	vm := evm.NewVM(data, contractState)
	vm.Run()

	fmt.Printf("%v\n", contractState)
	return data
}
