package main

import (
	"fmt"

	"github.com/TafveezA/EVM/evm"
)

func main() {
	contractdata := []byte{0x03, 0x0a, 0x02, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x046, 0x0c, 0x0d, 0x03, 0x0a, 0x0f}
	contractdata1 := []byte{}
	contractdata = append(contractdata, contractdata1...)
	contractState := evm.NewState()

	vm := evm.NewVM(contractdata, contractState)
	vm.Run()

	fmt.Printf("%v\n", contractState)
}
