package evm

type Instruction byte

const (
	InstrPush Instruction = 0x0a //10
	InstrAdd  Instruction = 0x0b //1
)

type Stack struct {
	data []any
	sp   int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]any, size),
		sp:   0,
	}

}

func (s *Stack) Push(v any) {
	s.data[s.sp] = v
	s.sp++
}

func (s *Stack) Pop() any {
	value := s.data[0]
	s.data = append(s.data[:0], s.data[1:]...)
	return value
}

type VM struct {
	data  []byte
	ip    int //instruction pointer
	stack []byte
	sp    int //stack pointer
}

func NewVM(data []byte) *VM {
	return &VM{
		data:  data,
		ip:    0,
		stack: make([]byte, 1024),
		sp:    -1,
	}
}

func (vm *VM) Run() error {
	for {
		instr := Instruction(vm.data[vm.ip])
		//fmt.Println(instr)
		if err := vm.Exec(instr); err != nil {
			return err
		}
		vm.ip++
		if vm.ip > len(vm.data)-1 {
			break
		}
	}
	return nil
}

func (vm *VM) Exec(instr Instruction) error {
	switch instr {
	case InstrPush:
		vm.pushStack(vm.data[vm.ip-1])
	case InstrAdd:
		a := vm.stack[0]
		b := vm.stack[1]
		c := a + b
		vm.pushStack(c)

	}
	return nil
}

func (vm *VM) pushStack(b byte) {
	vm.sp++
	vm.stack[vm.sp] = b

}
