package cpu

import (
	"github.com/Div9851/gba-emu/util"
)

const (
	EQ = iota
	NE
	CS
	CC
	MI
	PL
	VS
	VC
	HI
	LS
	GE
	LT
	GT
	LE
	AL
	NV
)

func (c *CPU) checkCond(opcode uint32) bool {
	cond := opcode >> 28

	switch cond {
	case EQ:
		return util.IsBitSet(c.CPSR, FlagZ)
	case NE:
		return !util.IsBitSet(c.CPSR, FlagZ)
	case CS:
		return util.IsBitSet(c.CPSR, FlagC)
	case CC:
		return !util.IsBitSet(c.CPSR, FlagC)
	case MI:
		return util.IsBitSet(c.CPSR, FlagN)
	case PL:
		return !util.IsBitSet(c.CPSR, FlagN)
	case VS:
		return util.IsBitSet(c.CPSR, FlagV)
	case VC:
		return !util.IsBitSet(c.CPSR, FlagV)
	case HI:
		return util.IsBitSet(c.CPSR, FlagC) && !util.IsBitSet(c.CPSR, FlagZ)
	case LS:
		return !util.IsBitSet(c.CPSR, FlagC) || util.IsBitSet(c.CPSR, FlagZ)
	case GE:
		return util.IsBitSet(c.CPSR, FlagN) == util.IsBitSet(c.CPSR, FlagV)
	case LT:
		return util.IsBitSet(c.CPSR, FlagN) != util.IsBitSet(c.CPSR, FlagV)
	case GT:
		return !util.IsBitSet(c.CPSR, FlagZ) && util.IsBitSet(c.CPSR, FlagN) == util.IsBitSet(c.CPSR, FlagV)
	case LE:
		return util.IsBitSet(c.CPSR, FlagZ) || util.IsBitSet(c.CPSR, FlagN) != util.IsBitSet(c.CPSR, FlagV)
	case AL:
		return true
	case NV:
		return false
	}

	panic("not reachable")
}
