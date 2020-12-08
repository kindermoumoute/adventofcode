package protocode

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

const (
	NOP      = "nop"
	ACC      = "acc"
	JMP      = "jmp"
	ADD      = "add"
	MUL      = "mul"
	IN       = "in"
	OUT      = "out"
	NON_ZERO = "ifnz"
	ZERO     = "ifz"
	LT       = "lt"
	EQ       = "eq"
	// don't forget to add new opcodes in knownOpcodes

	EOF = "eof"
)

var (
	knownOpcodes = map[string]struct {
		Word       string
		ParamCount int
	}{
		NOP:      {"NOP", 1},
		ACC:      {"ACC", 1},
		JMP:      {"JMP", 1},
		ADD:      {"ADD", 4},
		MUL:      {"MULTIPLY", 4},
		IN:       {"INPUT", 2},
		OUT:      {"OUTPUT", 2},
		EOF:      {"EOF", 1},
		NON_ZERO: {"JUMP IF NON ZERO", 3},
		ZERO:     {"JUMP IF ZERO", 3},
		LT:       {"IS LOWER THAN", 4},
		EQ:       {"IS EQUAL", 4},
	}
)

type ParameterMode int

const (
	ModePosition = ParameterMode(iota)
	ModeImmediate
	ModeRelative
)

type ProtoCode struct {
	Name                     string
	Memory                   map[int]*Instruction
	Cursor                   int
	IgnoreNonAddressedMemory bool
	DebugMode                DebugMode
	Acc                      int
	ExitCondition            func(c *ProtoCode) bool

	disablePostInc bool
}

type Instruction struct {
	Opcode string
	Params []int
	Used   int
}

// New copy the intcode memory before creating a new IntCode object.
func New(puzzle string) *ProtoCode {
	protoCode := &ProtoCode{
		Memory:                   map[int]*Instruction{},
		IgnoreNonAddressedMemory: true,
	}

	regLine := regexp.MustCompile(`([+-]?\d+)`)
	for i, line := range strings.Split(puzzle, "\n") {
		opcode := strings.Split(line, " ")
		matches := regLine.FindAllStringSubmatch(line, -1)
		protoCode.Memory[i] = &Instruction{
			Opcode: opcode[0],
		}
		for _, match := range matches {
			protoCode.Memory[i].Params = append(protoCode.Memory[i].Params, pkg.MustAtoi(match[1]))
		}
	}
	return protoCode
}

func (c *ProtoCode) Reset() {
	for _, instruction := range c.Memory {
		instruction.Used = 0
	}
	c.Acc = 0
	c.Cursor = 0
}

// Run the IntCode program
func (c *ProtoCode) Run() int {
	for {
		ins, exist := c.Memory[c.Cursor]
		if !exist || c.ExitCondition != nil && c.ExitCondition(c) {
			if c.Cursor != len(c.Memory) {
				c.infof("trying to read out of memory: %d", c.Cursor)
			}
			return c.Acc
		}
		ins.Used++
		c.printDebug(ins)

		switch ins.Opcode {
		case ACC:
			c.Acc += ins.Params[0]
			c.infof("acc=%d", c.Acc)
		case NOP:
			//c.Acc += ins.Params[0]
		case JMP: // relative jump for now
			c.Cursor += ins.Params[0]
			c.disablePostInc = true
		case ADD:
			//c.Write(3, c.Read(1)+c.Read(2))
		case MUL:
			//c.Write(3, c.Read(1)*c.Read(2))
		//case IN:
		//	c.Write(1, c.ReadInput())
		//case OUT:
		//	c.WriteOutput(c.Read(1))
		case NON_ZERO:
			//if c.Read(1) != 0 {
			//	c.Cursor = c.Read(2)
			//	c.disablePostInc = true
			//}
		case ZERO:
			//if c.Read(1) == 0 {
			//	c.Cursor = c.Read(2)
			//	c.disablePostInc = true
			//}
		case LT:
			//c.SetIf(3, c.Read(1) < c.Read(2))
		case EQ:
			//c.SetIf(3, c.Read(1) == c.Read(2))
		case EOF:
			return 0
		default:
			panic(fmt.Errorf("PROGRAM %s: %v unhandled", c.Name, ins.Opcode))
		}
		if !c.disablePostInc {
			c.Cursor++
		} else {
			c.disablePostInc = false
		}

	}
}

//
// Value reads the immediate value at address addr.
//func (c *ProtoCode) Value(addr, param int) int {
//	v, exist := c.Memory[addr]
//	if !exist && !c.IgnoreNonAddressedMemory {
//		//fmt.Println(c.Output)
//		panic(fmt.Errorf("error reading value at address %d: does not exist", addr))
//	}
//	return v.Params[param]
//}
//
//// Address reads the address value at position addr.
//func (c *ProtoCode) Address(addr int) int {
//	v, exist := c.Memory[c.Value(addr)]
//	if !exist && !c.IgnoreNonAddressedMemory {
//		//fmt.Println(c.Output)
//		panic(fmt.Errorf("error reading address at address %d from %d: does not exist", c.Value(addr), addr))
//	}
//	return v
//}
//
//// SetValue sets the given value into the address value at position addr.
//func (c *ProtoCode) SetValue(addr, value int) {
//	c.Memory[addr] = value
//}
//
//// SetAddress sets the given value into the address addr.
//func (c *ProtoCode) SetAddress(addr, value int) {
//	c.SetValue(c.Value(addr), value)
//}
//
//func (c *ProtoCode) SetIf(parameterMode ParameterMode, offset int, cond bool) {
//	if cond {
//		c.Write(parameterMode, offset, 1)
//	} else {
//		c.Write(parameterMode, offset, 0)
//	}
//}
//
//func (c *ProtoCode) Write(parameterMode ParameterMode, offset, value int) {
//	if parameterMode == ModeRelative {
//		c.SetValue(c.Value(c.Cursor+offset)+c.RelativeOffset, value)
//		return
//	}
//	c.SetAddress(c.Cursor+offset, value)
//}
//
//func (c *ProtoCode) Read(parameterMode ParameterMode, param int) int {
//	switch parameterMode {
//	case ModePosition:
//		return c.Address(c.Cursor + param)
//	case ModeRelative:
//		return c.Value(c.Value(c.Cursor, param))
//	case ModeImmediate:
//		return c.Value(c.Cursor, param)
//	}
//	panic(fmt.Errorf("unhandled mode %d", parameterMode))
//
//}
