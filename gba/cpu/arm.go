package cpu

import (
	"fmt"

	"github.com/Div9851/gba-emu/gba/context"
)

func (c *CPU) armStep(ctx *context.Context) {
	pc := c.R[15]
	c.pipe[1] = Inst{
		opcode: c.Bus.Read32(pc, true, ctx),
		addr:   pc,
	}
	c.armExec(c.inst.opcode, ctx)
	c.R[15] = pc + 4
}

func (c *CPU) armExec(opcode uint32, ctx *context.Context) {
	if !c.checkCond(opcode) {
		// If the condition is not met, do nothing.
		return
	}

	switch {
	case IsArmSWI(opcode):
		c.armSWI(opcode, ctx)
	case IsArmBL(opcode):
		c.armBL(opcode, ctx)
	case IsArmB(opcode):
		c.armB(opcode, ctx)
	case IsArmBX(opcode):
		c.armBX(opcode, ctx)
	case IsArmLDM(opcode):
		c.armLDM(opcode, ctx)
	case IsArmSTM(opcode):
		c.armSTM(opcode, ctx)
	case IsArmLDR(opcode):
		c.armLDR(opcode, ctx)
	case IsArmSTR(opcode):
		c.armSTR(opcode, ctx)
	case IsArmLDRH(opcode):
		c.armLDRH(opcode, ctx)
	case IsArmLDRSB(opcode):
		c.armLDRSB(opcode, ctx)
	case IsArmLDRSH(opcode):
		c.armLDRSH(opcode, ctx)
	case IsArmSTRH(opcode):
		c.armSTRH(opcode, ctx)
	case IsArmMRS(opcode):
		c.armMRS(opcode, ctx)
	case IsArmMSR(opcode):
		c.armMSR(opcode, ctx)
	case IsArmSWP(opcode):
		c.armSWP(opcode, ctx)
	case IsArmMPY(opcode):
		c.armMPY(opcode, ctx)
	case IsArmALU(opcode):
		c.armALU(opcode, ctx)
	default:
		panic(fmt.Sprintf("unknown opcode: 0x%08x", opcode))
	}

}

func (c *CPU) armB(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armBL(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armBX(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armSWI(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armALU(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armMPY(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armLDR(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armSTR(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armLDRH(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armLDRSB(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armLDRSH(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armSTRH(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armLDM(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armSTM(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armSWP(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armMRS(opcode uint32, ctx *context.Context) {

}

func (c *CPU) armMSR(opcode uint32, ctx *context.Context) {

}
