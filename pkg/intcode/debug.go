package intcode

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

func (c *IntCode) printDebug() {
	if c.DebugMode == DebugNone {
		return
	}
	if c.DebugMode == DebugVerbose {
		fmt.Println(c.String())
	}

	params := []string(nil)
	for i := 1; i < knownOpcodes[c.opcode].ParamCount; i++ {
		params = append(params, c.debugParam(c.Cursor+i, c.modes[i-1]))
	}
	c.infof("%s %s", knownOpcodes[c.opcode].Word, strings.Join(params, " "))
}

func (c *IntCode) debugParam(addr, mode int) string {
	if ParameterMode(mode) == ModeImmediate {
		return fmt.Sprintf("%d", c.Value(addr))
	}
	return fmt.Sprintf("%d[%d]", c.Address(addr), c.Value(addr))
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

func (c *IntCode) logDebug() {

}

func (c *IntCode) infof(s string, a ...interface{}) {
	if c.DebugMode != DebugNone {
		if c.Name != "" {
			s = "PROGRAM " + c.Name + ": " + s
		}
		fmt.Printf(strings.TrimRight(s, "\n")+"\n", a...)
	}
}
