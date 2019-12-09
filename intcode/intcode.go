package intcode

// OPCODES
const (
	Add        = 1
	Multiply   = 2
	Input      = 3
	Output     = 4
	JNZ        = 5
	JEZ        = 6
	LessThan   = 7
	Equals     = 8
	AdjustBase = 9
	Halt       = 99
)

const maxMemory = 2048

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
	RP     int
	ChIn   chan int
	ChOut  chan int
}

func (p *Program) LoadMemory(prog []int) {
	p.Memory = make([]int, maxMemory)
	copy(p.Memory, prog)
}

func (p *Program) Run() {
	p.IP = 0
	p.RP = 0

	for {
		opcode := parseOpcode(p.Memory[p.IP])
		switch opcode.code {
		case Add:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)
			_, dest := p.getArg(2, &opcode)

			p.Memory[dest] = arg1 + arg2
			p.IP += 4
		case Multiply:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)
			_, dest := p.getArg(2, &opcode)

			p.Memory[dest] = arg1 * arg2
			p.IP += 4
		case Input:
			_, dest := p.getArg(0, &opcode)
			value := <-p.ChIn
			p.Memory[dest] = value
			p.IP += 2
		case Output:
			arg1, _ := p.getArg(0, &opcode)
			p.ChOut <- arg1
			p.IP += 2
		case JNZ:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)

			if arg1 != 0 {
				p.IP = arg2
			} else {
				p.IP += 3
			}
		case JEZ:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)

			if arg1 == 0 {
				p.IP = arg2
			} else {
				p.IP += 3
			}
		case LessThan:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)
			_, dest := p.getArg(2, &opcode)

			if arg1 < arg2 {
				p.Memory[dest] = 1
			} else {
				p.Memory[dest] = 0
			}
			p.IP += 4
		case Equals:
			arg1, _ := p.getArg(0, &opcode)
			arg2, _ := p.getArg(1, &opcode)
			_, dest := p.getArg(2, &opcode)

			if arg1 == arg2 {
				p.Memory[dest] = 1
			} else {
				p.Memory[dest] = 0
			}

			p.IP += 4
		case AdjustBase:
			arg1, _ := p.getArg(0, &opcode)
			p.RP += arg1
			p.IP += 2
		case Halt:
			return
		default:
			panic("Invalid OpCode")
		}
	}
}

func (p *Program) getArg(i int, instr *opcode) (int, int) {
	switch instr.argModes[i] {
	case 0:
		return p.Memory[p.Memory[p.IP+1+i]], p.Memory[p.IP+1+i]
	case 1:
		return p.Memory[p.IP+1+i], 0
	case 2:
		return p.Memory[p.RP+(p.Memory[p.IP+1+i])], p.RP + (p.Memory[p.IP+1+i])
	}
	return 0, 0
}
