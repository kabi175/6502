package debugger

import (
	"fmt"

	"github.com/kabi175/6502/model"
)

func NewDebugger(bp []uint16, end uint16) model.Debugger {
	return new(Debugger).setBreakPoints(bp).setEnd(end)
}

type Debugger struct {
	bp  map[uint16]bool
	end uint16
}

func (d *Debugger) setBreakPoints(bp []uint16) *Debugger {
	for _, p := range bp {
		d.bp[p] = true
	}
	return d
}

func (d *Debugger) setEnd(end uint16) *Debugger {
	d.end = end
	return d
}

func (d *Debugger) IsEnd(pc uint16) bool {
	return d.end == pc
}

func (d *Debugger) Wait(pc uint16) {
	if ok := d.bp[pc]; !ok {
		return
	}
	var input string
	fmt.Scan(input)
}
