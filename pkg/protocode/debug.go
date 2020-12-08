package protocode

import (
	"fmt"
	"strings"
)

type DebugMode int

const (
	DebugNone = DebugMode(iota)
	DebugInfo
	DebugVerbose
)

func (c *ProtoCode) printDebug(ins *Instruction) {
	if c.DebugMode == DebugNone {
		return
	}
	if c.DebugMode == DebugVerbose {
		fmt.Println(c.String())
	}

	c.infof("%d: %s %v", c.Cursor, ins.Opcode, ins.Params)
}

func (c *ProtoCode) String() string {
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

func (c *ProtoCode) infof(s string, a ...interface{}) {
	if c.DebugMode != DebugNone {
		if c.Name != "" {
			s = "PROGRAM " + c.Name + ": " + s
		}
		fmt.Printf(strings.TrimRight(s, "\n")+"\n", a...)
	}
}
