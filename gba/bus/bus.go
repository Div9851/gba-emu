package bus

import "github.com/Div9851/gba-emu/gba/context"

type Bus struct{}

func New() *Bus {
	return &Bus{}
}

func (b *Bus) Read32(addr uint32, isSeq bool, ctx *context.Context) uint32 {
	return 0
}
