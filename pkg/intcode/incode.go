package intcode

import "fmt"

const (
	NOP = iota
	ADD_RRR
	MUL_RRR
	//SUB_RRR // <unofficial>
	//DIV_RRR
	//ADD_RVR
	//MUL_RVR
	//SUB_RVR
	//DIV_RVR
	//GT_RRR
	//GTE_RRR
	//EQ_RRR
	//LTE_RRR
	//LT_RRR // </unofficial>

	EOF = 99
)

var (
	knownOpCode = map[int]string{
		ADD_RRR: "adding",
		MUL_RRR: "multiplying",
		EOF:     "EOF",
	}
)

type IntCode struct {
	Memory map[int]int
	Cursor int

	IgnoreNonAddressedMemory bool
	Debug                    bool
}

func New(memory map[int]int, cursorStart int) *IntCode {
	intCode := &IntCode{
		Memory: map[int]int{},
		Cursor: cursorStart,
	}
	for k, v := range memory {
		intCode.Memory[k] = v
	}
	return intCode
}

func (c *IntCode) Run() *IntCode {
	for {
		target := c.Value(c.Cursor + 3) // debuging purpose
		if c.Debug {
			fmt.Println(c.String())
			fmt.Println(knownOpCode[c.Memory[c.Cursor]], c.Address(c.Cursor+1), c.Address(c.Cursor+2))
			fmt.Println("from addresses", c.Value(c.Cursor+1), c.Value(c.Cursor+2))
			fmt.Printf("to address %d currently at value %d\n", target, c.Value(target))
		}
		switch c.Memory[c.Cursor] {
		case ADD_RRR:
			c.SetAddress(c.Cursor+3, c.Address(c.Cursor+1)+c.Address(c.Cursor+2))
		case MUL_RRR:
			c.SetAddress(c.Cursor+3, c.Address(c.Cursor+1)*c.Address(c.Cursor+2))
		//case SUB_RRR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)-c.Address(c.Cursor+2))
		//case DIV_RRR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)/c.Address(c.Cursor+2))
		//case ADD_RVR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)+c.Value(c.Cursor+2))
		//case MUL_RVR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)*c.Value(c.Cursor+2))
		//case SUB_RVR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)-c.Value(c.Cursor+2))
		//case DIV_RVR:
		//	c.SetValue(c.Cursor+3,c.Address(c.Cursor+1)/c.Value(c.Cursor+2))
		//case GT_RRR:
		//	if c.Address(c.Cursor+1) > c.Address(c.Cursor+2) {
		//		c.Cursor = c.Address(c.Cursor+3)
		//	}
		//case GTE_RRR:
		//	if c.Address(c.Cursor+1) >= c.Address(c.Cursor+2) {
		//		c.Cursor = c.Address(c.Cursor+3)
		//	}
		//case EQ_RRR:
		//	if c.Address(c.Cursor+1) == c.Address(c.Cursor+2) {
		//		c.Cursor = c.Address(c.Cursor+3)
		//	}
		//case LTE_RRR:
		//	if c.Address(c.Cursor+1) <= c.Address(c.Cursor+2) {
		//		c.Cursor = c.Address(c.Cursor+3)
		//	}
		//case LT_RRR:
		//	if c.Address(c.Cursor+1) < c.Address(c.Cursor+2) {
		//		c.Cursor = c.Address(c.Cursor+3)
		//	}
		case EOF:
			return c
		default:
			panic("unhandled")
		}

		if c.Debug {
			fmt.Println("equal", c.Value(target))
		}
		c.Cursor += 4
	}
}

func (c *IntCode) Value(i int) int {
	v, exist := c.Memory[i]
	if !exist && !c.IgnoreNonAddressedMemory {
		panic(fmt.Errorf("error reading value at address %d: does not exist", i))
	}
	return v
}

func (c *IntCode) SetValue(addr, value int) {
	c.Memory[addr] = value
}

func (c *IntCode) SetAddress(addr, value int) {
	c.SetValue(c.Value(addr), value)
}

func (c *IntCode) Address(i int) int {
	v, exist := c.Memory[c.Value(i)]
	if !exist && !c.IgnoreNonAddressedMemory {
		panic(fmt.Errorf("error reading address at address %d: does not exist", i))
	}
	return v
}
func (c *IntCode) String() string {
	s := ""
	i := 0
	for {
		v, exist := c.Memory[i]
		if !exist {
			break
		}
		s += fmt.Sprint(v, ",")
		i++
	}
	s = s[:len(s)-1]
	s += fmt.Sprintf("\nCur: %d", c.Cursor)
	return s
}
