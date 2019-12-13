package intcode

import (
	"fmt"
	"math"

	"github.com/kindermoumoute/adventofcode/pkg"
)

const (
	NOP = iota
	ADD
	MUL
	IN
	OUT
	NON_ZERO
	ZERO
	LT
	EQ
	REL
	// don't forget to add new opcodes in knownOpcodes

	EOF = 99
)

var (
	knownOpcodes = map[int]struct {
		Word       string
		ParamCount int
	}{
		ADD:      {"ADD", 4},
		MUL:      {"MULTIPLY", 4},
		IN:       {"INPUT", 2},
		OUT:      {"OUTPUT", 2},
		EOF:      {"EOF", 1},
		NON_ZERO: {"JUMP IF NON ZERO", 3},
		ZERO:     {"JUMP IF ZERO", 3},
		LT:       {"IS LOWER THAN", 4},
		EQ:       {"IS EQUAL", 4},
		REL:      {"RELATIVE MODE", 2},
	}
)

type ParameterMode int

const (
	ModePosition = ParameterMode(iota)
	ModeImmediate
	ModeRelative
)

type IntCode struct {
	Name                     string
	Memory                   map[int]int
	Cursor                   int
	IgnoreNonAddressedMemory bool
	DebugMode                DebugMode
	Input                    ReadWriter
	Output                   ReadWriter
	OutputLinkedFromInput    bool
	Done                     chan bool

	disablePostInc bool
	opcode         int
	modes          []int
	RelativeOffset int
}

// ReadeWriter chan (C) expects a request and a response.
type ReadWriter struct {
	C    chan int
	Buff []int
}

// New copy the intcode memory before creating a new IntCode object.
func New(puzzle string, cursorStart int, seq ...int) *IntCode {
	intCode := &IntCode{
		Memory:                   map[int]int{},
		Cursor:                   cursorStart,
		IgnoreNonAddressedMemory: true,
		Input: ReadWriter{
			C:    make(chan int),
			Buff: seq,
		},
	}

	for k, v := range pkg.ParseIntMap(puzzle, ",") {
		intCode.Memory[k] = v
	}
	return intCode
}

// RunBackground runs in a go func.
func (c *IntCode) RunBackground() {
	go func() {
		c.Run()
	}()
}

// Run the IntCode program.
func (c *IntCode) Run() int {
	for {
		c.opcode, c.modes = readInstruction(c.Memory[c.Cursor])
		c.printDebug()

		switch c.opcode {
		case ADD:
			c.Write(3, c.Read(1)+c.Read(2))
		case MUL:
			c.Write(3, c.Read(1)*c.Read(2))
		case IN:
			c.Write(1, c.ReadInput())
		case OUT:
			c.WriteOutput(c.Read(1))
		case NON_ZERO:
			if c.Read(1) != 0 {
				c.Cursor = c.Read(2)
				c.disablePostInc = true
			}
		case ZERO:
			if c.Read(1) == 0 {
				c.Cursor = c.Read(2)
				c.disablePostInc = true
			}
		case LT:
			c.SetIf(3, c.Read(1) < c.Read(2))
		case EQ:
			c.SetIf(3, c.Read(1) == c.Read(2))
		case REL:
			c.RelativeOffset += c.Read(1)
		case EOF:
			if c.Done != nil {
				c.Done <- true
				close(c.Done)
			}
			if c.Output.C != nil && !c.OutputLinkedFromInput {
				close(c.Output.C)
			}
			if len(c.Output.Buff) > 0 {
				return c.Output.Buff[len(c.Output.Buff)-1]
			}
			return 0
		default:
			panic(fmt.Errorf("PROGRAM %s: %v unhandled", c.Name, c.opcode))
		}
		if !c.disablePostInc {
			c.Cursor += knownOpcodes[c.opcode].ParamCount
		} else {
			c.disablePostInc = false
		}
	}
}

func (c *IntCode) ReadInput() int {
	c.infof("\tstart reading")
	var in int
	if len(c.Input.Buff) > 0 { // read from input buffer first
		in = c.Input.Buff[0]
		c.Input.Buff = c.Input.Buff[1:]
	} else {

		if !c.OutputLinkedFromInput {
			c.infof("\trequest for input")
			c.Input.C <- 0 // request for input
			c.infof("\tinput request sent")
		}
		in = <-c.Input.C
	}
	c.infof("\treads %d", in)
	return in
}

func (c *IntCode) IntCodeToIntCodeProtocol(inputC chan int) {
	c.OutputLinkedFromInput = true
	c.Output.C = inputC
}

func (c *IntCode) WriteOutput(out int) {
	c.Output.Buff = append(c.Output.Buff, out)

	if c.Output.C != nil {
		c.infof("\tstart writing")
		if c.OutputLinkedFromInput {
			go func() {
				c.Output.C <- out
				c.infof("\twrote %d", out)
			}()
		} else {
			c.Output.C <- out
			c.infof("\twrote %d", out)

			<-c.Output.C // ack
			c.infof("\tack %d", out)
		}

		c.infof("\twrites %d", out)
	}
}

// Value reads the immediate value at address addr.
func (c *IntCode) Value(addr int) int {
	v, exist := c.Memory[addr]
	if !exist && !c.IgnoreNonAddressedMemory {
		fmt.Println(c.Output)
		panic(fmt.Errorf("error reading value at address %d: does not exist", addr))
	}
	return v
}

// Address reads the address value at position addr.
func (c *IntCode) Address(addr int) int {
	v, exist := c.Memory[c.Value(addr)]
	if !exist && !c.IgnoreNonAddressedMemory {
		fmt.Println(c.Output)
		panic(fmt.Errorf("error reading address at address %d from %d: does not exist", c.Value(addr), addr))
	}
	return v
}

// SetValue sets the given value into the address value at position addr.
func (c *IntCode) SetValue(addr, value int) {
	c.Memory[addr] = value
}

// SetAddress sets the given value into the address addr.
func (c *IntCode) SetAddress(addr, value int) {
	c.SetValue(c.Value(addr), value)
}

func (c *IntCode) SetIf(offset int, cond bool) {
	if cond {
		c.Write(offset, 1)
	} else {
		c.Write(offset, 0)
	}
}

func (c *IntCode) Write(offset, value int) {
	if ParameterMode(c.modes[offset-1]) == ModeRelative {
		c.SetValue(c.Value(c.Cursor+offset)+c.RelativeOffset, value)
		return
	}
	c.SetAddress(c.Cursor+offset, value)
}

func (c *IntCode) Read(offset int) int {
	switch ParameterMode(c.modes[offset-1]) {
	case ModePosition:
		return c.Address(c.Cursor + offset)
	case ModeRelative:
		return c.Value(c.Value(c.Cursor+offset) + c.RelativeOffset)
	case ModeImmediate:
		return c.Value(c.Cursor + offset)
	}
	panic(fmt.Errorf("unhandled mode %d", ParameterMode(c.modes[offset-1])))

}

func readInstruction(intructionCommand int) (opcode int, modes []int) {
	opcode = intructionCommand % 100
	for j := 100; j <= int(math.Pow10(knownOpcodes[opcode].ParamCount)); j *= 10 {
		modes = append(modes, (intructionCommand/j)%10)
	}
	return
}
