package cpu

import (
	"github.com/Div9851/gba-emu/gba/bus"
	"github.com/Div9851/gba-emu/gba/context"
	"github.com/Div9851/gba-emu/util"
)

type Inst struct {
	opcode uint32
	addr   uint32
}

const (
	FlagN = 31
	FlagZ = 30
	FlagC = 29
	FlagV = 28
	FlagT = 5
)

type CPU struct {
	Bus  *bus.Bus
	R    [16]uint32
	CPSR uint32
	// regBank [20]uint32
	inst Inst
	pipe []Inst
}

func New() *CPU {
	return &CPU{
		Bus:  bus.New(),
		pipe: make([]Inst, 2),
	}
}

func (c *CPU) IsArmState() bool {
	return util.IsBitSet(c.CPSR, FlagT)
}

// Execute the next instruction.
func (c *CPU) Step(ctx *context.Context) {
	c.inst = c.pipe[0]
	c.pipe[0] = c.pipe[1]

	if c.IsArmState() {
		c.armStep(ctx)
	} else {
		c.thumbStep(ctx)
	}
}
