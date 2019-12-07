package intcode

// OPCODES
const (
	Add      = 1
	Multiply = 2
	Input    = 3
	Output   = 4
	JNZ      = 5
	JEZ      = 6
	LessThan = 7
	Equals   = 8
	Halt     = 99
)

type opcode struct {
	code     int
	argModes []int
}

func parseOpcode(code int) opcode {
	opcode := opcode{}
	digit := 0

	digit = code % 10
	code = code / 10
	opcode.code = ((code % 10) * 10) + digit
	code = code / 10

	for i := 0; i < 3; i++ {
		opcode.argModes = append(opcode.argModes, code%10)
		code = code / 10
	}

	return opcode
}

type Program struct {
	Memory []int
	IP     int
	ChIn   chan int
	ChOut  chan int
}

func (p *Program) Run() {
	p.IP = 0

	for {
		opcode := parseOpcode(p.Memory[p.IP])
		switch opcode.code {
		case Add:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)
			dest := p.Memory[p.IP+3]

			p.Memory[dest] = arg1 + arg2
			p.IP += 4
		case Multiply:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)
			dest := p.Memory[p.IP+3]

			p.Memory[dest] = arg1 * arg2
			p.IP += 4
		case Input:
			dest := p.Memory[p.IP+1]
			value := <-p.ChIn
			p.Memory[dest] = value
			p.IP += 2
		case Output:
			arg1 := p.getArg(0, &opcode)
			p.ChOut <- arg1
			p.IP += 2
		case JNZ:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)

			if arg1 != 0 {
				p.IP = arg2
			} else {
				p.IP += 3
			}
		case JEZ:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)

			if arg1 == 0 {
				p.IP = arg2
			} else {
				p.IP += 3
			}
		case LessThan:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)
			dest := p.Memory[p.IP+3]

			if arg1 < arg2 {
				p.Memory[dest] = 1
			} else {
				p.Memory[dest] = 0
			}
			p.IP += 4
		case Equals:
			arg1 := p.getArg(0, &opcode)
			arg2 := p.getArg(1, &opcode)
			dest := p.Memory[p.IP+3]

			if arg1 == arg2 {
				p.Memory[dest] = 1
			} else {
				p.Memory[dest] = 0
			}

			p.IP += 4
		case Halt:
			return
		default:
			panic("Invalid OpCode")
		}
	}
}

func (p *Program) getArg(i int, instr *opcode) int {
	if instr.argModes[i] == 1 {
		return p.Memory[p.IP+1+i]
	}
	return p.Memory[p.Memory[p.IP+1+i]]
}
